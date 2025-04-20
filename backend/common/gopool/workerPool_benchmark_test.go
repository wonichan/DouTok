package gopool

import (
	"sync"
	"testing"
	"time"
)

// 模拟任务，可以设置执行时间
type benchmarkJob struct {
	executionTime time.Duration
	id            int
}

func (j *benchmarkJob) Do() error {
	if j.executionTime > 0 {
		time.Sleep(j.executionTime)
	}
	return nil
}

// 基准测试：测试不同数量的worker和任务下的性能
func BenchmarkWorkerPool(b *testing.B) {
	tests := []struct {
		name       string
		maxWorkers uint16
		workerNum  uint16
		jobTime    time.Duration
	}{
		{"Small_Pool_Fast_Jobs", 10, 5, time.Microsecond * 100},
		{"Small_Pool_Slow_Jobs", 10, 5, time.Millisecond},
		{"Medium_Pool_Fast_Jobs", 50, 25, time.Microsecond * 100},
		{"Medium_Pool_Slow_Jobs", 50, 25, time.Millisecond},
		{"Large_Pool_Fast_Jobs", 200, 100, time.Microsecond * 100},
		{"Large_Pool_Slow_Jobs", 200, 100, time.Millisecond},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			pool := NewWorkerPool(tt.maxWorkers, tt.workerNum)
			defer pool.Close()

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				job := &benchmarkJob{executionTime: tt.jobTime, id: i}
				pool.Accept(job)
			}
		})
	}
}

// 基准测试：测试与普通goroutine方式的对比
func BenchmarkWorkerPoolVsGoroutines(b *testing.B) {
	tests := []struct {
		name       string
		maxWorkers uint16
		workerNum  uint16
		jobTime    time.Duration
		jobCount   int
	}{
		{"Small_Workload", 50, 25, time.Microsecond * 100, 30},
		{"Medium_Workload", 100, 50, time.Microsecond * 100, 60},
		{"Large_Workload", 200, 100, time.Microsecond * 100, 150},
	}

	for _, tt := range tests {
		// 测试WorkerPool
		b.Run("WorkerPool_"+tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				pool := NewWorkerPool(tt.maxWorkers, tt.workerNum)
				b.StartTimer()

				for j := 0; j < tt.jobCount; j++ {
					job := &benchmarkJob{executionTime: tt.jobTime, id: j}
					pool.Accept(job)
				}

				b.StopTimer()
				pool.Close()
				b.StartTimer()
			}
		})

		// 测试普通goroutine
		b.Run("Goroutines_"+tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				var wg sync.WaitGroup
				b.StartTimer()

				for j := 0; j < tt.jobCount; j++ {
					wg.Add(1)
					go func(id int) {
						defer wg.Done()
						job := &benchmarkJob{executionTime: tt.jobTime, id: id}
						job.Do()
					}(j)
				}

				b.StopTimer()
				wg.Wait()
				b.StartTimer()
			}
		})
	}
}

// 基准测试：测试高并发场景下的性能
func BenchmarkWorkerPoolHighConcurrency(b *testing.B) {
	tests := []struct {
		name       string
		maxWorkers uint16
		workerNum  uint16
		concurrent int
	}{
		{"Low_Concurrency", 100, 50, 10},
		{"Medium_Concurrency", 100, 50, 50},
		{"High_Concurrency", 100, 50, 100},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			pool := NewWorkerPool(tt.maxWorkers, tt.workerNum)
			defer pool.Close()

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					job := &benchmarkJob{executionTime: time.Microsecond * 100}
					pool.Accept(job)
				}
			})
		})
	}
}

// 基准测试：测试动态调整worker数量的性能影响
func BenchmarkWorkerPoolAdjustment(b *testing.B) {
	tests := []struct {
		name           string
		initialWorkers uint16
		adjustTo       uint16
		jobCount       int
	}{
		{"Scale_Up", 10, 100, 1000},
		{"Scale_Down", 100, 10, 1000},
		{"No_Change", 50, 50, 1000},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				pool := NewWorkerPool(200, tt.initialWorkers)
				b.StartTimer()

				// 提交一半任务
				halfJobs := tt.jobCount / 2
				for j := 0; j < halfJobs; j++ {
					job := &benchmarkJob{executionTime: time.Microsecond * 50}
					pool.Accept(job)
				}

				// 调整worker数量
				pool.AdjustWorkerNum(tt.adjustTo)

				// 提交剩余任务
				for j := 0; j < halfJobs; j++ {
					job := &benchmarkJob{executionTime: time.Microsecond * 50}
					pool.Accept(job)
				}

				b.StopTimer()
				pool.Close()
				b.StartTimer()
			}
		})
	}
}
