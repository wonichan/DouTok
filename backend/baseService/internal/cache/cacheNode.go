package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dgraph-io/ristretto"

	"DouTok-example/backend/baseService/internal/cache/types"
	"DouTok-example/backend/baseService/internal/logger"
)

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
	sync.RWMutex // 缓存未命中时获取源数据的回调函数
	delayTTL     time.Duration
	cache        *ristretto.Cache
	stats        NodeStats
}

func NewCacheNode(cacheSize int, delayTTL time.Duration) *CacheNode {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: int64(cacheSize * 10), // 估计的键数量的10倍
		MaxCost:     int64(cacheSize),      // 最大容量
		BufferItems: 64,                    // 缓冲区大小
		OnEvict: func(item *ristretto.Item) {
			fmt.Printf("Evicted item with key: %d, value:%+v\n", item.Key, item.Value)
		},
		OnExit: func(val interface{}) {
			fmt.Printf("item exit: %+v\n", val)
		},
	})
	if err != nil {
		panic(err)
	}
	return &CacheNode{
		delayTTL: delayTTL,
		cache:    cache,
	}
}

// 从缓存中获取数据
func (cn *CacheNode) Get(key string) (kv *types.KvStore, ok bool) {
	// 从其他节点获取
	defer func() {
		if kv != nil {
			if kv.IsExpired() {
				cn.Del(kv.Key)
				kv = nil
				ok = false
			}
		}
	}()
	// 从本地缓存获取
	if kvCache, ok := cn.cache.Get(key); ok {
		logger.GetLogger().Infof("load cache from lru cache")
		atomic.AddUint64(&cn.stats.HitCount, 1)
		kv = kvCache.(*types.KvStore)
		return kv, ok
	}
	// 缓存击穿
	atomic.AddUint64(&cn.stats.MissCount, 1)
	return kv, ok
}

func (cn *CacheNode) Set(key string, value *types.KvStore) {
	logger.GetLogger().Infof("set cache to lru cache, ttl:%v", value.LifeTime+cn.delayTTL)
	_ = cn.cache.SetWithTTL(key, value, value.Length(), value.LifeTime+cn.delayTTL)
	cn.stats.WriteCount++
}

func (cn *CacheNode) Del(key string) {
	cn.cache.Del(key)
}

func (cn *CacheNode) PopulateCache(kv *types.KvStore) {
	cn.cache.SetWithTTL(kv.Key, kv, kv.Length(), kv.LifeTime+cn.delayTTL)
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
		// 其他字段...
	}
}
