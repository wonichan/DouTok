package cache

import (
	"sync"

	"DouTok-example/backend/baseService/internal/cache/consistenthash"
	"DouTok-example/backend/baseService/internal/cache/types"
)

const (
	defaultReplicas = 50
	self            = "node1"
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

func (p *CacheNodesPool) AddNodes(addrs ...string) *CacheNodesPool {
	p.hashRing.Add(addrs...)
	for _, addr := range addrs {
		p.Nodes[addr] = NewPeerClient(addr)
	}
	return p
}

func (p *CacheNodesPool) PickPeer(key string) (types.PeerGetter, bool) {
	p.RLock()
	defer p.RUnlock()
	if addr := p.hashRing.Get(key); addr != "" && addr != self {
		if peer, ok := p.Nodes[addr]; ok {
			return peer, true
		}
	}
	return nil, false
}

type PeerClient struct {
	Addr string
}

func NewPeerClient(addr string) *PeerClient {
	return &PeerClient{
		Addr: addr,
	}
}

func (p *PeerClient) Get(group, key string) ([]byte, error) {
	return []byte{}, nil
}
