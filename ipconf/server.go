package ipconf

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunMain() {
	h := server.Default(server.WithHostPorts(":4000"))
	h.GET("/ip/list", GetIpInfoList)
	h.Spin()
}
