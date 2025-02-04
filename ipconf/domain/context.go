package domain

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// IpConfConext is the context of the ipconf module.
type IpConfContext struct {
	Ctx       *context.Context
	AppReqCtx *app.RequestContext
}

// BuildIpConfContext creates a new IpConfConext.
func BuildIpConfContext(c *context.Context, ctx *app.RequestContext) *IpConfContext {
	ipConfContext := &IpConfContext{
		Ctx:       c,
		AppReqCtx: ctx,
	}
	return ipConfContext
}
