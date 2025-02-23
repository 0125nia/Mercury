package discovery

import (
	"context"
	"testing"
	"time"

	"github.com/0125nia/Mercury/common/config"
)

func TestServiceDiscovery(t *testing.T) {
	config.InitConfig("./../../mercury.yaml")
	ctx := context.Background()
	ser := NewServiceDiscovery(&ctx)
	defer ser.Close()
	ser.WatchService("/web/", func(key, value string) {}, func(key, value string) {})
	ser.WatchService("/gRPC/", func(key, value string) {}, func(key, value string) {})

	for range time.Tick(10 * time.Second) {
	}
}
