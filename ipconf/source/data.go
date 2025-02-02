package source

import (
	"context"

	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/ipconf/discovery"
)

func Init() {
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
	discovery := discovery.NewServiceDiscovery(ctx)
	defer discovery.Close()
	// todo add data handler here
}
