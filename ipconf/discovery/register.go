package discovery

import (
	"context"
	"log"
	"time"

	"github.com/0125nia/Mercury/common/config"
	"github.com/bytedance/gopkg/util/logger"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ServiceRegister Lease Registration Service
type ServiceRegister struct {
	cli           *clientv3.Client                        // etcd client
	leaseID       clientv3.LeaseID                        // lease id
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // keep alive channel
	key           string                                  // etcd key
	value         string                                  // etcd value
	ctx           *context.Context                        // context
}

// NewServiceRegister creates a new service register
func NewServiceRegister(ctx *context.Context, key string, endpoint *Endpoint, lease int64) (*ServiceRegister, error) {
	// create a new etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Config.Discovery.Endpoints,
		DialTimeout: config.Config.Discovery.TimeOut * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// create service register
	register := &ServiceRegister{
		cli:   cli,
		key:   key,
		value: endpoint.Marshal(),
		ctx:   ctx,
	}

	// apply for lease and set lease time
	if err := register.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return register, nil
}

// putKeyWithLease register the service with lease
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	// set the lease time
	resp, err := s.cli.Grant(*s.ctx, lease)
	if err != nil {
		return err
	}
	// register the service and bind the lease
	_, err = s.cli.Put(*s.ctx, s.key, s.value, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}

	// keep the lease alive
	keepAliveChannel, err := s.cli.KeepAlive(*s.ctx, resp.ID)
	if err != nil {
		return err
	}
	s.leaseID = resp.ID
	s.keepAliveChan = keepAliveChannel
	return nil
}

// UpdateValue update the value of the service
func (s *ServiceRegister) UpdateValue(val *Endpoint) error {
	// Marshal the Endpoint
	value := val.Marshal()
	_, err := s.cli.Put(*s.ctx, s.key, s.value, clientv3.WithLease(s.leaseID))
	if err != nil {
		return err
	}
	s.value = value
	logger.CtxInfof(*s.ctx, "ServiceRegister.updateValue leaseID=%d Put key=%s,val=%s, success!", s.leaseID, s.key, s.value)
	return nil
}

// keepAlive is a goroutine to keep the lease alive
func (s *ServiceRegister) KeepAlive() {
	for leaseKeepResp := range s.keepAliveChan {
		logger.CtxInfof(*s.ctx, "lease success leaseID:%d, Put key:%s,val:%s reps:+%v",
			s.leaseID, s.key, s.value, leaseKeepResp)
	}
	logger.CtxInfof(*s.ctx, "lease failed !!!  leaseID:%d, Put key:%s,val:%s", s.leaseID, s.key, s.value)
}

// Close the lease
func (s *ServiceRegister) Close() error {
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	logger.CtxInfof(*s.ctx, "lease Close... leaseID:%d key:%s,value:%s", s.leaseID, s.key, s.value)
	return s.cli.Close()
}
