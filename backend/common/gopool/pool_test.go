package gopool

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/stretchr/testify/assert"
)

// TestNewPool tests the creation of a new pool with default and custom options.
func TestNewPool(t *testing.T) {
	// Test with default options
	pool, err := NewPool()
	assert.NoError(t, err, "NewPool with default options should not return an error")
	assert.NotNil(t, pool, "NewPool with default options should return a non-nil pool")
	assert.NotNil(t, pool.antsPool, "Internal ants pool should be initialized")
	assert.Equal(t, ants.DefaultAntsPoolSize, pool.Cap(), "Default pool capacity should match ants default")
	pool.Release()

	// Test with custom MaxSize
	customSize := 10
	pool, err = NewPool(WithMaxSize(customSize))
	assert.NoError(t, err, "NewPool with custom MaxSize should not return an error")
	assert.NotNil(t, pool, "NewPool with custom MaxSize should return a non-nil pool")
	assert.Equal(t, customSize, pool.Cap(), "Pool capacity should match custom MaxSize")
	pool.Release()

	// Test with invalid MaxSize (handled by ants internally, should return error)
	_, err = NewPool(WithMaxSize(-1), WithAntsOptions(ants.WithPreAlloc(true)))
	assert.Error(t, err, "NewPool with negative MaxSize should return an error")
}

// TestPoolSubmit tests basic task submission.
func TestPoolSubmit(t *testing.T) {
	pool, err := NewPool(WithMaxSize(5))
	assert.NoError(t, err)
	defer pool.Release()

	var counter int32
	var wg sync.WaitGroup
	numTasks := 10
	wg.Add(numTasks)

	taskFunc := func() {
		atomic.AddInt32(&counter, 1)
		wg.Done()
	}

	for i := 0; i < numTasks; i++ {
		err = pool.Submit(Task(taskFunc))
		assert.NoError(t, err, "Submit should not return an error for valid tasks")
	}

	wg.Wait() // Wait for all tasks to complete
	assert.Equal(t, int32(numTasks), atomic.LoadInt32(&counter), "All submitted tasks should have executed")
}

// TestPoolSubmitTimeout tests task submission with timeout.
func TestPoolSubmitTimeout(t *testing.T) {
	pool, err := NewPool(WithMaxSize(1), WithSubmitTimeout(20*time.Millisecond), WithAntsOptions(ants.WithNonblocking(true)))
	assert.NoError(t, err)
	defer pool.Release()

	var wg sync.WaitGroup
	wg.Add(1)

	// Submit a task that blocks the only worker
	blockingTask := func() {
		time.Sleep(100 * time.Millisecond) // Block longer than timeout
		wg.Done()
	}
	err = pool.Submit(Task(blockingTask))
	assert.NoError(t, err)

	// Try to submit another task, which should time out
	quickTask := func() { t.Log("This should not run") }
	err = pool.Submit(Task(quickTask))
	assert.ErrorIs(t, err, ErrTimeout, "Submitting to a full pool with timeout should return ErrTimeout")

	wg.Wait() // Wait for the blocking task to finish
}

// TestPoolRelease tests releasing pool resources.
func TestPoolRelease(t *testing.T) {
	pool, err := NewPool(WithMaxSize(5))
	assert.NoError(t, err)

	var counter int32
	numTasks := 3
	// Submit some tasks
	for i := 0; i < numTasks; i++ {
		err = pool.Submit(Task(func() {
			atomic.AddInt32(&counter, 1)
			time.Sleep(50 * time.Millisecond) // Simulate work
		}))
		assert.NoError(t, err)
	}

	// Give tasks a moment to start
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, numTasks, pool.Running(), "Running count should reflect submitted tasks")

	pool.Release()

	// Allow time for running tasks to potentially finish after Release() is called
	time.Sleep(100 * time.Millisecond)

	// Attempt to submit after release (should error)
	err = pool.Submit(Task(func() { t.Log("Should not run after release") }))
	assert.Error(t, err, "Submit should fail after pool is released")

	// Check if pool stats reflect release
	assert.Equal(t, 5, pool.Cap(), "Capacity should remain after release")
	// Note: Ants doesn't guarantee Running() is 0 immediately after Release(),
	// especially if tasks were still running.
}

// TestPoolStats tests the pool status functions.
func TestPoolStats(t *testing.T) {
	poolSize := 3
	pool, err := NewPool(WithMaxSize(poolSize))
	assert.NoError(t, err)
	defer pool.Release()

	assert.Equal(t, poolSize, pool.Cap(), "Cap() should return the configured pool size")
	assert.Equal(t, 0, pool.Running(), "Running() should be 0 initially")
	assert.Equal(t, poolSize, pool.Free(), "Free() should be equal to Cap() initially")

	var wg sync.WaitGroup
	numToSubmit := 2
	wg.Add(numToSubmit)

	// Submit fewer tasks than pool size
	for i := 0; i < numToSubmit; i++ {
		err = pool.Submit(Task(func() {
			time.Sleep(50 * time.Millisecond) // Keep worker busy
			wg.Done()
		}))
		assert.NoError(t, err)
	}

	// Allow some time for tasks to start
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, poolSize, pool.Cap(), "Cap() should remain constant")
	assert.Equal(t, numToSubmit, pool.Running(), "Running() should reflect active tasks")
	assert.Equal(t, poolSize-numToSubmit, pool.Free(), "Free() should reflect available workers")

	wg.Wait() // Wait for tasks to complete

	// Allow some time for workers to become free
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, poolSize, pool.Cap(), "Cap() should remain constant after tasks finish")
	assert.Equal(t, numToSubmit, pool.Running(), "Running() should be 0 after tasks finish")
	assert.Equal(t, poolSize-numToSubmit, pool.Free(), "Free() should return to Cap() after tasks finish")
}

// BenchmarkPool benchmarks the performance of the custom Pool.
func BenchmarkPool(b *testing.B) {
	pool, _ := NewPool(WithMaxSize(1000)) // Use a larger size for benchmark
	defer pool.Release()

	var wg sync.WaitGroup
	task := func() {
		// Simulate minimal work
		_ = 1 + 1
		wg.Done()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		_ = pool.Submit(Task(task))
	}
	wg.Wait()
}

// BenchmarkGoroutine benchmarks the performance of standard Go goroutines.
func BenchmarkGoroutine(b *testing.B) {
	var wg sync.WaitGroup
	task := func() {
		// Simulate minimal work
		_ = 1 + 1
		wg.Done()
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go task()
	}
	wg.Wait()
}
