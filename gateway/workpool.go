package gateway

import (
	"fmt"

	"github.com/0125nia/Mercury/common/config"
	"github.com/panjf2000/ants"
)

var wPool *ants.Pool

func initWorkPool() {
	var err error
	if wPool, err = ants.NewPool(config.Config.Gateway.WorkerPoolNum); err != nil {
		fmt.Printf("InitWorkPoll.err :%s num:%d\n", err.Error(), config.Config.Gateway.WorkerPoolNum)
	}
}
