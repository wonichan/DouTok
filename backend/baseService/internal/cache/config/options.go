package config

import (
	"DouTok-example/backend/baseService/internal/cache/types"
	"time"
)

const (
	DefaultCacheSize = 1000
	DefaultTTL       = time.Minute * 5
	DefualtDelayTTL  = time.Minute
)

var (
	DefaultGetter = types.GetterFunc(func(key string) (*types.KvStore, error) {
		return &types.KvStore{}, nil
	})
)

type Options struct {
	CacheSize  int
	DefaultTTL time.Duration
	DelayTTL   time.Duration
	Getter     types.Getter
}

type Option struct {
	F func(*Options)
}

func NewOptions(op ...Option) *Options {
	opts := &Options{
		DefaultTTL: DefaultTTL,
		CacheSize:  DefaultCacheSize,
		DelayTTL:   DefualtDelayTTL,
		Getter:     DefaultGetter,
	}
	opts.Apply(op...)
	return opts
}

func (o *Options) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.F(o)
	}
}

func WithDefaultTTL(ttl time.Duration) Option {
	return Option{
		F: func(o *Options) {
			o.DefaultTTL = ttl
		},
	}
}

func WithDefaultDelayTTL(ttl time.Duration) Option {
	return Option{
		F: func(o *Options) {
			o.DelayTTL = ttl
		},
	}
}
func WithCacheSize(size int) Option {
	return Option{
		F: func(o *Options) {
			o.CacheSize = size
		},
	}
}

func WithGetter(getter types.Getter) Option {
	return Option{
		F: func(o *Options) {
			o.Getter = getter
		},
	}
}
