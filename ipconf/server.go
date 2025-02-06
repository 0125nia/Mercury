package ipconf

import (
	"github.com/0125nia/Mercury/common/config"
	"github.com/0125nia/Mercury/ipconf/domain"
	"github.com/0125nia/Mercury/ipconf/source"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunMain(path string) {
	config.InitConfig(path)
	// Initialize data source
	source.Init()
	// Initialize the dispatch layer
	domain.Init()
	// Start the web server
	h := server.Default(server.WithHostPorts(":4000"))
	h.GET("/ip/list", GetIpInfoList)
	h.Spin()
}
