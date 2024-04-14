package config

import (
	"path/filepath"

	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

var Props *Config

type Config struct {
	Listen string
}

func Parse(location string) bool {
	if location == "" {
		location = filepath.Join(location, "conf.yaml")
	}

	viper.SetConfigFile(location)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		golog.Errorf("config file %s read failed. %v\n", location, err)
	}
	err = viper.GetViper().Unmarshal(&Props)
	if err != nil {
		golog.Errorf("sfu config file %s loaded failed. %v\n", location, err)
	}
	return true
}
