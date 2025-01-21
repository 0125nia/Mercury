package ipconf

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// IpConfConext is the context of the ipconf module.
type IpConfConext struct {
	Ctx       *context.Context
	AppReqCtx *app.RequestContext
}

// BuildIpConfContext creates a new IpConfConext.
func BuildIpConfContext(c *context.Context, ctx *app.RequestContext) *IpConfConext {
	ipConfConext := &IpConfConext{
		Ctx:       c,
		AppReqCtx: ctx,
	}
	return ipConfConext
}
