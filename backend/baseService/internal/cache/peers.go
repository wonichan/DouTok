package cache

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"

	"DouTok-example/backend/baseService/api"
	"DouTok-example/backend/baseService/api/cacheservice"
	"DouTok-example/backend/baseService/internal/cache/consistenthash"
	"DouTok-example/backend/baseService/internal/cache/types"
	"DouTok-example/backend/baseService/internal/logger"
)

const (
	defaultReplicas = 50
)

type CacheNodesPool struct {
	sync.RWMutex
	hashRing *consistenthash.Map
	Nodes    map[string]types.PeerGetter
}

func NewCacheNodesPool() *CacheNodesPool {
	return &CacheNodesPool{
		hashRing: consistenthash.New(defaultReplicas, nil),
		Nodes:    make(map[string]types.PeerGetter),
	}
}

func (p *CacheNodesPool) AddNodes(addrs ...string) {
	p.hashRing.Add(addrs...)
	for _, addr := range addrs {
		p.Nodes[addr] = NewPeerClient(addr)
	}
}

func (p *CacheNodesPool) PickPeer(key string) (types.PeerGetter, bool) {
	p.RLock()
	defer p.RUnlock()
	if addr := p.hashRing.Get(key); addr != "" && addr != selfNode {
		if peer, ok := p.Nodes[addr]; ok {
			return peer, true
		}
	}
	return nil, false
}

func (p *CacheNodesPool) GetKV(ctx context.Context, req *api.CacheRequest) (res *api.CacheResponse, err error) {
	logger.GetLogger().Infof("get kv from aother node")
	kv, ok := GetCacheGroup(req.Group).Get(req.Key)
	err = errors.New("failed to get value")
	if !ok {
		res.Value = nil
		err = errors.New("failed to get value")
		return
	}
	res = &api.CacheResponse{}
	res.Value, err = json.Marshal(kv)
	return
}

func (p *CacheNodesPool) startRPCServer() error {
	conn, err := net.ResolveTCPAddr("tcp", selfNode)
	if err != nil {
		return err
	}
	srv := cacheservice.NewServer(p,
		server.WithServiceAddr(conn),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "CacheService"}),
	)
	go func() {
		if err := srv.Run(); err != nil {
			panic(err)
		}
	}()
	return nil
}

type PeerClient struct {
	Addr   string
	client cacheservice.Client
}

func NewPeerClient(addr string) *PeerClient {
	fmt.Printf("new peer client: %s\n", addr)
	return &PeerClient{
		Addr: addr,
		client: cacheservice.MustNewClient("CacheService",
			client.WithHostPorts(addr),
			client.WithMuxConnection(100),
			client.WithTransportProtocol(transport.GRPC),
		),
	}
}

func (p *PeerClient) Get(group, key string) ([]byte, error) {
	logger.GetLogger().Infof("get kv from aother node:%s", p.Addr)
	r, err := p.client.GetKV(context.Background(), &api.CacheRequest{
		Group: group,
		Key:   key,
	})
	if err != nil {
		return nil, err
	}
	return r.Value, nil
}
