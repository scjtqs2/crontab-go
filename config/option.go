package config

import (
	"github.com/go-yaml/yaml"
	"github.com/scjtqs/utils/util"
)

type Conf struct {
	Crontab []Crontab
}

type Crontab struct {
	Cron string `yaml:"cron"`
	Cmd  string `yaml:"command"`
}

// 通过路径获取配置信息
func GetConfigFronPath(c string) *Conf {
	conf := &Conf{}
	if !util.PathExists(c) {
		conf = defaultConf()
	} else {
		err := yaml.Unmarshal([]byte(util.ReadAllText(c)), conf)
		if err != nil {
			conf = defaultConf()
		}
	}
	return parseConfFromEnv(conf)
}

// 没有配置文件时候的默认配置
func defaultConf() *Conf {
	return &Conf{
		[]Crontab{
			{
				Cron: "*/10 * * * * *",
				Cmd:  "echo 'hello world'",
			},
			{
				Cron: "0 * * * *",
				Cmd:  "echo 'hello !!!'",
			},
		},
	}
}

// 从环境变量中替换配置文件
func parseConfFromEnv(c *Conf) *Conf {
	//todu
	return c
}

// 保存配置文件
func (c *Conf) Save(p string) error {
	s, _ := yaml.Marshal(c)
	return util.WriteAllText(p, string(s))
}
