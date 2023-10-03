package workbench_service

import (
	"time"

	"github.com/integralist/go-elasticache/elasticache"

	aws "github.com/sheikhrachel/workbench/api_common/aws_utils"
	"github.com/sheikhrachel/workbench/api_common/call"
)

type WorkbenchService struct {
	aws             *aws.AWSClient
	memcachedClient *elasticache.Client
}

func New(cc call.Call, awsClient *aws.AWSClient) *WorkbenchService {
	service := &WorkbenchService{aws: awsClient}
	service.setupMemcached(cc)
	return service
}

const (
	itemTTL int32 = 300
)

// setupMemcached sets up the memcached client by reading the ELASTICACHE_ENDPOINT environment variable
func (s *WorkbenchService) setupMemcached(cc call.Call) {
	var err error
	done := make(chan bool)
	go func() {
		s.memcachedClient, err = elasticache.New()
		done <- true
	}()
	select {
	case <-done:
		if err != nil {
			return
		}
		cc.InfoF("Successfully set up memcached client")
	case <-time.After(3 * time.Second):
	}
}
