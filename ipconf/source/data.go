package source

import (
	"context"

	"github.com/0125nia/Mercury/common/config"
)

func Init() {
	ctx := context.Background()
	go DataHandler(&ctx)
	if config.Config.Global.Env == "debug" {
		// todo add debug test here
	}
}

func DataHandler(ctx *context.Context) {

}
