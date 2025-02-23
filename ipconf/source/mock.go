package source

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/common/discovery"
)

// testServiceRegister is a function to test the service register
func testServiceRegister(ctx *context.Context, port, node string) {
	// mock test the service register
	go func() {
		ed := discovery.EndpointInfo{
			Ip:   "127.0.0.1",
			Port: port,
			MetaData: map[string]interface{}{
				"connect_num":   float64(rand.Int63n(12312321231231131)),
				"message_bytes": float64(rand.Int63n(1231232131556)),
			},
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
			ed = discovery.EndpointInfo{
				Ip:   ed.Ip,
				Port: ed.Port,
				MetaData: map[string]interface{}{
					"connect_num":   float64(rand.Int63n(12312321231231131)),
					"message_bytes": float64(rand.Int63n(1231232131556)),
				},
			}
			sr.UpdateValue(&ed)
			time.Sleep(1 * time.Second)
		}
	}()
}
