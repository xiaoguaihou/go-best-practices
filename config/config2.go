package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfigViper struct {
	isInit bool
}

func (config *AppConfigViper) GetString(key string) string {
	return viper.GetString(key)
}

func (config *AppConfigViper) GetInt(key string) int {
	return viper.GetInt(key)
}

func (config *AppConfigViper) GetConfig(file string, isForce ...bool) {

	force := false

	if len(isForce) == 1 {
		force = isForce[0]
	}

	if !config.isInit || force {
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err == nil {
			config.isInit = true
		} else {
			log.Fatalf("failed to parse configure file [%s], err:%+v\n", file, err)
		}
	}
}
