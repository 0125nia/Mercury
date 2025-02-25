package gateway

import (
	"fmt"
	"log"
	"net"

	"github.com/0125nia/Mercury/common/config"
)

func RunMain(configPath string) {
	config.InitConfig(configPath)

	_, err := net.ListenTCP("tcp", &net.TCPAddr{Port: config.Config.Gateway.ServerPort})
	if err != nil {
		log.Panicf("Start epoll server failed: %v", err)
	}

	initWorkPool()

	// todo init epoll

	fmt.Println("Start epoll server success!")

	select {}
}
