package ipconf

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetIpInfoList(c context.Context, ctx *app.RequestContext) {
	defer func() {
		// Recover from panic
		if r := recover(); r != nil {
			ctx.JSON(consts.StatusBadRequest, utils.H{"err": r})
		}
	}()

	// Build the IpConfContext
	ipConfCtx := BuildIpConfContext(&c, ctx)

	// ip dispatch processing
	dispatch(ipConfCtx)

	// resp the ip info list
	ctx.JSON(consts.StatusOK, ipConfResp())
}

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
