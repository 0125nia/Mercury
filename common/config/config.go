package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	Global    *Global    `yaml:"global"`
	Discovery *Discovery `yaml:"discovery"`
	IpConf    *IpConf    `yaml:"ipconf"`
	// todo add config here
}

type Global struct {
	Env string `yaml:"env"`
	// todo add global config here
}

type Discovery struct {
	Endpoints []string      `yaml:"endpoints"`
	TimeOut   time.Duration `yaml:"timeout"`
}

type IpConf struct {
	ServicePath string `yaml:"service_path"`
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
