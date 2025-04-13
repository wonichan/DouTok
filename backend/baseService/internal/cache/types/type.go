package types

type Getter interface {
	Get(key string) (*KvStore, error)
}

type GetterFunc func(key string) (*KvStore, error)

func (f GetterFunc) Get(key string) (*KvStore, error) {
	return f(key)
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
