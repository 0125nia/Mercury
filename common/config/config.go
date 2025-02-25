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
	Gateway   *Gateway   `yaml:"gateway"`
}

type Global struct {
	Env string `yaml:"env"`
}

type Discovery struct {
	Endpoints []string      `yaml:"endpoints"`
	TimeOut   time.Duration `yaml:"timeout"`
}

type IpConf struct {
	ServicePath string `yaml:"service_path"`
}

type Gateway struct {
	WorkerPoolNum  int   `yaml:"workerpool_num"`
	ServerPort     int   `yaml:"server_port"`
	EpollerChanNum int   `yaml:"epoller_channel_num"`
	EpollerNum     int   `yaml:"epoller_num"`
	TCPMaxNum      int32 `yaml:"tcp_max_num"`
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
