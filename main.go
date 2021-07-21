package main

import (
	"flag"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/scjtqs/crontab-go/app"
	"github.com/scjtqs/crontab-go/config"
	"github.com/scjtqs/utils/util"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"path"
	"time"
)

var (
	h          bool
	d          bool
	Version    = "v1.0.0"
	Build      string
	configPath = "config.yml"
)

func init() {
	var debug bool
	flag.BoolVar(&d, "d", false, "running as a daemon")
	//flag.BoolVar(&debug, "D", false, "debug mode")
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&configPath, "c", "config.yml", "config file path default is config.yml")
	flag.Parse()
	logFormatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	}
	w, err := rotatelogs.New(path.Join("logs", "%Y-%m-%d.log"), rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}
	LogLevel := "info"
	if debug {
		log.SetReportCaller(true)
		LogLevel = "debug"
	}
	log.AddHook(util.NewLocalHook(w, logFormatter, util.GetLogLevel(LogLevel)...))
}

func main() {
	if h {
		help()
	}
	if d {
		util.Daemon()
	}
	conf := config.GetConfigFronPath(configPath)
	container := dig.New()
	container.Provide(func() *config.Conf {
		return conf
	})
	conf.Save(configPath)
	log.Infof("welcome to use crontab-go  by scjtqs  https://github.com/scjtqs/crontab-go %s,build in %s", Version, Build)
	app.Start(container)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

// help cli命令行-h的帮助提示
func help() {
	log.Infof(`crontab service
version: %s
built-on: %s

Usage:

server [OPTIONS]

Options:
`, Version, Build)
	flag.PrintDefaults()
	os.Exit(0)
}