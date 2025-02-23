package perf

import (
	"net"

	"github.com/0125nia/Mercury/common/sdk"
)

var (
	TcpConnNum int32
)

func RunMain() {
	for i := 0; i < int(TcpConnNum); i++ {
		sdk.NewChat(net.ParseIP("127.0.0.1"), 8900, "mercury", "12", "111")
	}
}
