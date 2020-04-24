package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	DB struct {
		SalesHost string `yaml:"salesHost"`
		SalesUser string `yaml:"salesUser"`
		SalesPwd  string `yaml:"salesPwd"`

		OrderHost string `yaml:"orderHost"`
		OrderUser string `yaml:"orderUser"`
		OrderPwd  string `yaml:"orderPwd"`

		PoolSize int `yaml:"poolSize"`
	} `yaml:"DB"`

	RocketMq string `yaml:"rocketmq"`
	Timer    int    `yaml:"timer"`

	Port   int `yaml:"port"`
	isInit bool
}

func (config *AppConfig) GetConfig(file string, isForce ...bool) {

	force := false

	if len(isForce) == 1 {
		force = isForce[0]
	}

	if !config.isInit || force {
		data, err := ioutil.ReadFile(file)
		if err == nil {
			err := yaml.Unmarshal(data, config)
			if err != nil {
				log.Fatalf("failed to parse configure file [%s], err:%+v\n", file, err)
			}
			config.isInit = true
		} else {
			log.Fatalf("failed to open configure file [%s], err:%+v\n", file, err)
		}
	}
}
