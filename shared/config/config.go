package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	imOnce sync.Once
	conf   *Config
)

type Config struct {
	Port     int      `mapstructure:"PORT"`
	DATABASE DATABASE `mapstructure:"DATABASE"`
	POSTGRES POSTGRES `mapstructure:"POSTGRES"`
	MONGO    MONGO    `mapstructure:"MONGO"`
}

type DATABASE struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Host     string `mapstructure:"HOST"`
	Name     string `mapstructure:"NAME"`
}

type POSTGRES struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Name     string `mapstructure:"NAME"`
}

type MONGO struct {
	URI string `mapstructure:"URI"`
}

func GetDefaultImmutableConfig() *Config {
	var outer error
	var success = true

	env := os.Getenv("APP_ENV")
	pwd := os.Getenv("APP_PWD")

	if env == "test" && pwd == "" {
		panic(fmt.Errorf("APP_PWD env is required in test env"))
	}

	imOnce.Do(func() {
		v := viper.New()
		v.SetConfigName("app.config")

		if env == "test" {
			v.AddConfigPath(pwd)
		} else {
			v.AddConfigPath(".")
		}

		v.SetEnvPrefix("vp")
		v.AutomaticEnv()

		if err := v.ReadInConfig(); err != nil {
			success = false
			outer = fmt.Errorf("failed to read app.config file due to %s", err)
		}

		v.Unmarshal(&conf)
	})

	if !success {
		panic(outer)
	}

	return conf
}
