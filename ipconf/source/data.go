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
		// todo add debug test here
	}
}

// DataHandler is a goroutine to handle the service discovering
func DataHandler(ctx *context.Context) {
	d := discovery.NewServiceDiscovery(ctx)
	defer d.Close()

	// setFunc is a function to handle the new service
	setFunc := func(key, value string) {
		if ed, err := discovery.UnmarshalEndpoint([]byte(value)); err == nil {
			if event := NewEvent(AddNodeEvent, ed); ed != nil {
				eventChan <- event
			}
		} else {
			logger.CtxErrorf(*ctx, "DataHandler.setFunc.err :%s", err.Error())
		}
	}

	// delFunc is a function to handle the deleted service
	delFunc := func(key, value string) {
		if ed, err := discovery.UnmarshalEndpoint([]byte(value)); err == nil {
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
