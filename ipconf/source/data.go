package source

import (
	"context"

	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/ipconf/discovery"
	"github.com/bytedance/gopkg/util/logger"
)

func Init() {
	// init event channel
	eventChan = make(chan *Event)

	ctx := context.Background()
	// goroutine to handle data
	go DataHandler(&ctx)
	// debug test
	if config.Config.Global.Env == "debug" {
		ctx := context.Background()
		testServiceRegister(&ctx, "8081", "node1")
		testServiceRegister(&ctx, "8082", "node2")
		testServiceRegister(&ctx, "8083", "node3")
	}
}

// DataHandler is a goroutine to handle the service discovering
func DataHandler(ctx *context.Context) {
	d := discovery.NewServiceDiscovery(ctx)
	defer d.Close()

	// setFunc is a function to handle the new service
	setFunc := func(key, value string) {
		if ed, err := discovery.UnmarshalEndpointInfo([]byte(value)); err == nil {
			if event := NewEvent(AddNodeEvent, ed); ed != nil {
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
	}

	// delFunc is a function to handle the deleted service
	delFunc := func(key, value string) {
		if ed, err := discovery.UnmarshalEndpointInfo([]byte(value)); err == nil {
			if event := NewEvent(DelNodeEvent, ed); ed != nil {
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.delFunc.err :%s", err.Error())
		}
	}

	err := d.WatchService(config.Config.IpConf.ServicePath, setFunc, delFunc)
	if err != nil {
		panic(err)
	}
}
