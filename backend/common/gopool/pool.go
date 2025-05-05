package gopool

import (
	"errors"
	"time"

	"github.com/panjf2000/ants/v2"
)

// ErrTimeout is returned when a task submission times out.
var ErrTimeout = errors.New("task submission timed out")

type TaskInterface interface {
	Do()
}

// Task represents a function to be executed by the pool.
type Task func()

func (t Task) Do() {
	t()
}

// Pool is a generic goroutine pool based on ants.
type Pool struct {
	antsPool *ants.Pool
	options  Options
}

// Options holds the configuration for the pool.
type Options struct {
	// MaxSize is the maximum number of goroutines in the pool.
	MaxSize int
	// SubmitTimeout is the maximum duration to wait for submitting a task.
	// If zero or negative, there is no timeout.
	SubmitTimeout time.Duration
	// AntsOptions allows passing custom options to the underlying ants pool.
	AntsOptions []ants.Option
}

// Option applies a configuration option to the pool.
type Option func(*Options)

// WithMaxSize sets the maximum size of the pool.
func WithMaxSize(size int) Option {
	return func(o *Options) {
		o.MaxSize = size
	}
}

// WithSubmitTimeout sets the maximum duration to wait for submitting a task.
func WithSubmitTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.SubmitTimeout = timeout
	}
}

// WithAntsOptions sets custom options for the underlying ants pool.
func WithAntsOptions(opts ...ants.Option) Option {
	return func(o *Options) {
		o.AntsOptions = append(o.AntsOptions, opts...)
	}
}

// NewPool creates a new goroutine pool.
func NewPool(opts ...Option) (*Pool, error) {
	options := Options{
		MaxSize: ants.DefaultAntsPoolSize, // Default size
	}
	for _, opt := range opts {
		opt(&options)
	}

	pool, err := ants.NewPool(options.MaxSize, options.AntsOptions...)
	if err != nil {
		return nil, err
	}

	return &Pool{
		antsPool: pool,
		options:  options,
	}, nil
}

// Submit submits a task to the pool.
// If a SubmitTimeout is configured and the submission takes longer, ErrTimeout is returned.
func (p *Pool) Submit(task TaskInterface) error {
	if p.options.SubmitTimeout <= 0 {
		return p.antsPool.Submit(task.Do)
	}

	timer := time.NewTimer(p.options.SubmitTimeout)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			return ErrTimeout
		default:
			// Try to submit without blocking indefinitely
			err := p.antsPool.Submit(task.Do)
			// Handle potential blocking scenario if Submit itself could block
			// Note: ants.Submit might block if Nonblocking is false and pool is full.
			// Consider using Submit non-blockingly or handling this case.
			if err == ants.ErrPoolOverload {
				time.Sleep(1 * time.Millisecond) // Short sleep to avoid busy waiting
				continue
			} else if err != nil {
				return err
			}
		}
	}
}

// Release stops the pool and releases resources.
func (p *Pool) Release() {
	p.antsPool.Release()
}

// Running returns the number of running goroutines.
func (p *Pool) Running() int {
	return p.antsPool.Running()
}

// Cap returns the capacity of the pool.
func (p *Pool) Cap() int {
	return p.antsPool.Cap()
}

// Free returns the number of available goroutines.
func (p *Pool) Free() int {
	return p.antsPool.Free()
}
