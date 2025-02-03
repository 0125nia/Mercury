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
	// mock test the service register
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

		// keep alive channel handle
		go sr.KeepAlive()

		// mock the service update
		for {
			ed = discovery.Endpoint{
				Ip:   ed.Ip,
				Port: ed.Port,
			}
			sr.UpdateValue(&ed)
			time.Sleep(1 * time.Second)
		}
	}()
}
