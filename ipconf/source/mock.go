package source

import (
	"context"
	"fmt"
	"time"

	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/ipconf/discovery"
)

// testServiceRegister is a function to test the service register
func testServiceRegister(ctx *context.Context, port, node string) {
	go func() {
		ed := discovery.Endpoint{
			Ip:   "127.0.0.1",
			Port: port,
		}
		// register the service
		sr, err := discovery.NewServiceRegister(ctx, fmt.Sprintf("%s/%s", config.Config.IpConf.ServicePath, node), &ed, time.Now().Unix())
		if err != nil {
			panic(err)
		}
		go sr.KeepAlive()
		// todo: update endpoint here
	}()
}
