package cache

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"

	"DouTok-example/backend/baseService/internal/cache/singleflight"
	"DouTok-example/backend/baseService/internal/cache/types"
	"DouTok-example/backend/baseService/internal/logger"
)

var DefaultGetter = types.GetterFunc(func(key string) (types.ByteView, bool) {
	return types.ByteView{}, false
})

type NodeStats struct {
	StartTime  time.Time
	MemoryUsed uint64
	ItemCount  uint64
	HitCount   uint64 // 原子计数器
	MissCount  uint64
	EvictCount uint64
	WriteCount uint64
	WriteQPS   uint64
	ReadQPS    uint64
}

type CacheNode struct {
	sync.RWMutex
	Addr         string
	virtualNodes uint64
	getter       types.Getter                               // 缓存未命中时获取源数据的回调函数
	loader       singleflight.Group[string, types.ByteView] // 防止缓存击穿，从其他节点获取数据
	stats        NodeStats
	lruCache     *lru.Cache[string, types.ByteView]
}

func NewCacheNode(addr string, virtualNodes uint64, cacheSize int, getter types.Getter) *CacheNode {
	lruCache, _ := lru.New[string, types.ByteView](cacheSize)
	if getter == nil {
		getter = DefaultGetter
	}
	return &CacheNode{
		Addr:         addr,
		virtualNodes: virtualNodes,
		lruCache:     lruCache,
		getter:       getter,
	}
}

func (cn *CacheNode) Get(key string) (types.ByteView, bool) {
	b, ok := cn.lruCache.Get(key)
	logger.GetLogger().Infof("get cache from %s", cn.Addr)
	if ok {
		logger.GetLogger().Infof("load cache from lru cache")
		atomic.AddUint64(&cn.stats.HitCount, 1)
		return b, ok
	}
	// 缓存击穿
	atomic.AddUint64(&cn.stats.MissCount, 1)

	if b, ok = cn.GetWithLoader(key); ok {
		logger.GetLogger().Infof("load cache from other node")
	} else {
		// 缓存未命中，从getter获取
		logger.GetLogger().Warnf("load cache from getter")
		b, ok = cn.getter.Get(key)
	}
	if ok {
		cn.lruCache.Add(key, b)
		cn.stats.WriteCount++
	}
	return b, ok
}

func (cn *CacheNode) GetWithLoader(key string) (types.ByteView, bool) {
	b, _, err := cn.loader.Do(context.Background(), key, func(ctx context.Context) (types.ByteView, error) {
		// 从其他节点获取数据
		b, errFromOtherNode := cn.GetFromOtherNodes(key)
		if errFromOtherNode != nil {
			return types.ByteView{}, errFromOtherNode
		}
		return b, nil
	})
	if err != nil {
		// 处理错误
		return types.ByteView{}, false
	}
	return b, true
}

// TODO: 从其他节点获取数据
func (cn *CacheNode) GetFromOtherNodes(key string) (types.ByteView, error) {
	// 从其他节点获取数据
	return types.ByteView{}, errors.New("not implemented") // 暂时返回错误
}

func (cn *CacheNode) Set(key string, value types.ByteView) {
	logger.GetLogger().Infof("set cache from %s", cn.Addr)
	_ = cn.lruCache.Add(key, value)
	cn.stats.WriteCount++
}

func (cn *CacheNode) Del(key string) {
	cn.lruCache.Remove(key)
}

// 统计信息上报
func (n *CacheNode) ReportStats() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	/* for range ticker.C {
		stats := n.snapshotStats()
		cluster.metricsCollector.Push(n.nodeID, stats)

		// 自动触发扩容
		if stats.MemoryUsed > n.config.HighWatermark {
			cluster.ScaleOut(1)
		}
	} */
}

// 生成统计快照（避免锁竞争）
func (cn *CacheNode) snapshotStats() NodeStats {
	cn.RLock()
	defer cn.RUnlock()

	return NodeStats{
		MemoryUsed: atomic.LoadUint64(&cn.stats.MemoryUsed),
		ItemCount:  uint64(cn.lruCache.Len()),
		// 其他字段...
	}
}
