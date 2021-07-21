package app

import (
	"github.com/robfig/cron/v3"
	"github.com/scjtqs/crontab-go/config"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

type cronSvr struct {
	Ct   *dig.Container
	Conf *config.Conf
	Cron *cron.Cron
}

func Start(ct *dig.Container) {
	svr := cronSvr{
		Ct: ct,
	}
	ct.Invoke(func(cf *config.Conf) {
		svr.Conf = cf
	})

	svr.Cron = cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))

	for k, v := range svr.Conf.Crontab {
		crontab := v.Cron
		cmd := v.Cmd
		log.Infof("index %d cron: %s  command: %s registed", (k + 1), crontab, cmd)
		// 添加回调
		_, err := svr.Cron.AddFunc(crontab, func() {
			log.Warnf("crontab cron %s command: %s start", crontab, cmd)
			//执行cmd
			var proc *exec.Cmd
			switch runtime.GOOS {
			case "windows":
				proc = exec.Command("cmd", "/C", cmd)
				break
			case "drawin":
				proc = exec.Command("bash", "-c", cmd)
				break
			default:
				proc = exec.Command("sh", "-c", cmd)
				break
			}
			res, err := proc.Output()
			if err != nil {
				log.Errorf("crontab cron %s command: %s faild,error=%s", crontab, cmd, err.Error())
				return
			}
			log.Infof("crontab cron %s command: %s successed,res=%s", crontab, cmd, string(res))
		})
		if err != nil {
			panic("cron start with error:" + err.Error())
		}
	}
	svr.Cron.Start()

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		svr.Cron.Stop()
		os.Exit(1)
	}()
}
