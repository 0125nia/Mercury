package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	Global *Global `yaml:"global"`
	// todo add config here
}

type Global struct {
	Env string `yaml:"env"`
	// todo add global config here
}

func InitConfig(path string) {
	workDir, _ := os.Getwd()
	viper.SetConfigName(path)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
