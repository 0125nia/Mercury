package ipconf

import (
	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/ipconf/source"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunMain(path string) {
	config.InitConfig(path)
	source.Init()
	h := server.Default(server.WithHostPorts(":4000"))
	h.GET("/ip/list", GetIpInfoList)
	h.Spin()
}
