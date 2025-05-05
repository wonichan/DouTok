package types

import (
	"time"

	"DouTok-example/backend/common/redis"
)

type Getter interface {
	Get(key string) (*KvStore, error)
}

type GetterFunc func(key string) (*KvStore, error)

func (f GetterFunc) Get(key string) (*KvStore, error) {
	return f(key)
}

type DCS struct {
	redisClient *redis.RedisClient
}

func (dcs *DCS) Get(key string) (*KvStore, error) {
	byte, err := dcs.redisClient.GetBytes(key)
	if err != nil {
		return nil, err
	}
	kv := &KvStore{
		Key:        key,
		Value:      byte,
		LifeTime:   0,
		UpdateTime: time.Now(),
	}
	if ttl, err := dcs.redisClient.TTL(key); err != nil {
		return nil, err
	} else {
		kv.LifeTime = ttl
	}

	return kv, nil
}

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Group string `json:"group"`
	TTL   int64  `json:"ttl"`
}
