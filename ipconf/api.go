package ipconf

import (
	"context"

	"github.com/0125nia/Mercury/ipconf/domain"
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
	ipConfCtx := domain.BuildIpConfContext(&c, ctx)

	// ip dispatch processing
	eds := domain.Dispatch(ipConfCtx)

	// resp the ip info list
	ipConfCtx.AppReqCtx.JSON(consts.StatusOK, ipConfResp(top5Endpoints(eds)))
}
