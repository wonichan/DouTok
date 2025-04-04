package types

type Getter interface {
	Get(key string) (ByteView, bool)
}

type GetterFunc func(key string) (ByteView, bool)

func (f GetterFunc) Get(key string) (ByteView, bool) {
	return f(key)
}

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
