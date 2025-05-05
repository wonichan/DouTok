package types

import (
	"time"
)

// 创建key/value模型
type KvStore struct {
	Key        string        `json:"key"`
	Value      ByteView      `json:"value"`
	LifeTime   time.Duration `json:"lifeTime"` // 单位秒
	UpdateTime time.Time     `json:"updateTime"`
}

// 创建一条key/value内容
func NewKvStore(key string, value ByteView, lifeTime time.Duration) *KvStore {
	return &KvStore{
		key,
		value,
		lifeTime,
		time.Now()}
}

// 检验是否过期
func (kvstore *KvStore) IsExpired() bool {
	//判断永不过期
	if kvstore.LifeTime == 0 {
		return false
	}
	return time.Since(kvstore.UpdateTime) > kvstore.LifeTime
}

func (kvstore *KvStore) Length() int64 {
	return int64(kvstore.Value.Len())
}

func (kvstore *KvStore) Clone() *KvStore {
	return &KvStore{
		Key:        kvstore.Key,
		Value:      kvstore.Value.ByteSlice(),
		LifeTime:   kvstore.LifeTime,
		UpdateTime: time.Now(),
	}
}
