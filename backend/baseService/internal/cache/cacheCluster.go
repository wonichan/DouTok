package cache

import (
	"sync"

	"DouTok-example/backend/baseService/internal/cache/consistenthash"
)

const defaultReplicas = 50

type CacheCluster struct {
	sync.RWMutex
	hashRing *consistenthash.Map
	nodes    map[string]*CacheNode
}

func NewCacheCluster() *CacheCluster {
	return &CacheCluster{
		hashRing: consistenthash.New(defaultReplicas, nil),
		nodes:    make(map[string]*CacheNode),
	}
}

func (cc *CacheCluster) AddNodes(cacheNodes ...*CacheNode) {
	cc.Lock()
	defer cc.Unlock()
	for _, n := range cacheNodes {
		cc.hashRing.Add(n.Addr)
		cc.nodes[n.Addr] = n
	}
}

func (cc *CacheCluster) GetNode(key string) *CacheNode {
	cc.RLock()
	defer cc.RUnlock()

	return cc.nodes[cc.hashRing.Get(key)]
}

func (cc *CacheCluster) RemoveNode(node *CacheNode) {
	cc.Lock()
	defer cc.Unlock()
	cc.hashRing.Remove(node.Addr)
	delete(cc.nodes, node.Addr)
}
