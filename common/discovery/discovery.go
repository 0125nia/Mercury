package discovery

import (
	"context"
	"sync"
	"time"

	"github.com/0125nia/Mercury/common/config"
	"github.com/bytedance/gopkg/util/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ServiceDiscovery is a struct of service discovering
type ServiceDiscovery struct {
	etcdClient *clientv3.Client
	lock       sync.Mutex
	ctx        *context.Context
}

// NewServiceDiscovery creates a new service discovery
func NewServiceDiscovery(ctx *context.Context) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Config.Discovery.Endpoints,
		DialTimeout: config.Config.Discovery.TimeOut * time.Second,
	})

	if err != nil {
		logger.Fatal(err)
	}
	return &ServiceDiscovery{
		etcdClient: cli,
		ctx:        ctx,
	}
}

// WatchService init the service list and watches the service
func (s *ServiceDiscovery) WatchService(prefix string, set, del func(key, value string)) error {

	// get the key with the prefix
	resp, err := s.etcdClient.Get(*s.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	// init the service list
	for _, kv := range resp.Kvs {
		set(string(kv.Key), string(kv.Value))
	}

	// watch the prefix of the key and update the server which is updated
	// resp.Header.Revision + 1 is the revision of the key
	// the reason why plus 1 is making sure the data is more recent than the current revision, avoid working on the same revisions repeatedly
	s.watcher(prefix, resp.Header.Revision+1, set, del)

	return nil
}

// watcher watches the service
func (s *ServiceDiscovery) watcher(prefix string, rev int64, set, del func(key, value string)) {
	// Get the channel for watching the status of services with the prefix
	watchChan := s.etcdClient.Watch(*s.ctx, prefix, clientv3.WithPrefix(), clientv3.WithRev(rev))
	logger.CtxInfof(*s.ctx, "watching prefix: %s ...", prefix)

	for WatchResp := range watchChan {
		for _, event := range WatchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut: // update or add
				set(string(event.Kv.Key), string(event.Kv.Value))
			case clientv3.EventTypeDelete: // delete
				del(string(event.Kv.Key), string(event.Kv.Value))
			}
		}
	}
}

// Close closes the service discovery
func (s *ServiceDiscovery) Close() error {
	return s.etcdClient.Close()
}
