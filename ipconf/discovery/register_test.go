package discovery

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/0125nia/Mercury/common/config"
)

func TestServciceRegister(t *testing.T) {
	config.InitConfig("./../../mercury.yaml")
	ctx := context.Background()
	ser, err := NewServiceRegister(&ctx, "/web/node1", &Endpoint{
		Ip:   "127.0.0.1",
		Port: "9999",
	}, 5)
	if err != nil {
		log.Fatalln(err)
	}

	go ser.KeepAlive()
	time.Sleep(20 * time.Second)
	ser.Close()
}
