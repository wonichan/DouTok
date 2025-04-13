package cache

import (
	"DouTok-example/backend/baseService/internal/cache/config"
	"DouTok-example/backend/baseService/internal/cache/singleflight"
	"DouTok-example/backend/baseService/internal/cache/types"
	"DouTok-example/backend/baseService/internal/logger"
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/json"
)

var (
	mu     sync.RWMutex
	groups = make(map[string]*CacheGroup)
)

func NewCacheGroupPool(name string, opts ...config.Option) {
	mu.Lock()
	defer mu.Unlock()
	groups[name] = NewCacheGroup(name, opts...)
}

func GetCacheGroup(name string) *CacheGroup {
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}

type CacheGroup struct {
	name      string
	delayTTL  time.Duration
	mainCache *CacheNode
	getter    types.Getter
	nodesPool *CacheNodesPool
	loader    *singleflight.Group[string, *types.KvStore]
}

func NewCacheGroup(name string, opts ...config.Option) *CacheGroup {
	options := config.NewOptions(opts...)
	return &CacheGroup{
		name:      name,
		mainCache: NewCacheNode(options.CacheSize, options.DelayTTL),
		getter:    options.Getter,
		loader:    &singleflight.Group[string, *types.KvStore]{},
	}
}

func (cg *CacheGroup) RegisterNode(nodePool *CacheNodesPool) {
	if nodePool == nil {
		panic("nodePool is nil")
	}
	cg.nodesPool = nodePool
}

func (cg *CacheGroup) Get(key string) (kv *types.KvStore, ok bool) {
	// 先从本地缓存获取
	if kv, ok = cg.mainCache.Get(key); ok {
		logger.GetLogger().Infof("load cache from local cache")
		return kv, ok
	}
	// 从其他节点获取或从本地获取
	logger.GetLogger().Infof("load cache from other node or local")
	return cg.load(key)
}

func (cg *CacheGroup) load(key string) (*types.KvStore, bool) {
	// 从其他节点获取
	kv, _, err := cg.loader.Do(context.Background(), key, func(_ context.Context) (*types.KvStore, error) {
		if cacheNode, ok := cg.nodesPool.PickPeer(key); ok {
			if otherKV, otherErr := cg.getFromPeer(cacheNode, key); otherErr == nil {
				logger.GetLogger().Infof("load cache from other node")
				return otherKV, nil
			}
		}
		// 从本地获取
		logger.GetLogger().Infof("load cache from local")
		return cg.getLocally(key)
	})
	if err != nil {
		return nil, false
	}
	return kv, true
}

func (cg *CacheGroup) getLocally(key string) (*types.KvStore, error) {
	kv, err := cg.getter.Get(key)
	if err != nil {
		return nil, err
	}
	// 从持久化存储设备获取时，是否考虑加上延迟时间
	if kv.LifeTime != 0 {
		kv.LifeTime = kv.LifeTime + cg.delayTTL
	}
	cg.PopulateCache(kv.Clone())
	return kv, nil
}

func (cg *CacheGroup) getFromPeer(peer types.PeerGetter, key string) (*types.KvStore, error) {
	byte, err := peer.Get(cg.name, key)
	if err != nil {
		logger.GetLogger().Errorf("get cache from peer failed: %v", err)
		return nil, err
	}
	kv := &types.KvStore{}
	if err := json.Unmarshal(byte, kv); err != nil {
		return nil, err
	}
	return kv, nil
}

func (cg *CacheGroup) PopulateCache(kv *types.KvStore) {
	cg.mainCache.PopulateCache(kv)
}

func (cg *CacheGroup) Set(kv *types.KvStore) {
	cg.mainCache.Set(kv.Key, kv)
}
