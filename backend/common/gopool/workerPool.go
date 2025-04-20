package gopool

import (
	"errors"
	"fmt"
	"sync"
)

type WorkerPool struct {
	closed       bool
	maxWorkers   uint16
	aliveWorkers uint16
	workerNum    uint16
	mux          sync.RWMutex
	closeOnce    sync.Once
	onPanic      func(interface{})
	workerQueue  chan *Worker
}

func NewWorkerPool(maxWorkers, workerNum uint16) *WorkerPool {
	if workerNum > maxWorkers {
		workerNum = maxWorkers
	}
	return &WorkerPool{
		maxWorkers:  maxWorkers,
		workerNum:   workerNum,
		mux:         sync.RWMutex{},
		closeOnce:   sync.Once{},
		workerQueue: make(chan *Worker, maxWorkers),
	}
}

func (wp *WorkerPool) OnPanic(fn func(interface{})) {
	wp.onPanic = fn
}

func (wp *WorkerPool) Pause() {
	wp.AdjustWorkerNum(0)
}

func (wp *WorkerPool) Resume(num uint16) {
	wp.AdjustWorkerNum(num)
}

// 调整worker数量
func (wp *WorkerPool) AdjustWorkerNum(num uint16) {
	wp.mux.Lock()
	defer wp.mux.Unlock()
	if wp.closed {
		return
	}
	if num > wp.maxWorkers {
		num = wp.maxWorkers
	}
	oldNum := wp.workerNum
	wp.workerNum = num
	if oldNum > num {
		// 减少worker数量
		for num < wp.aliveWorkers {
			worker := <-wp.workerQueue
			worker.Close()
			wp.aliveWorkers--
		}
	} else if oldNum < num {
		// 增加worker数量
		for num > wp.aliveWorkers {
			wp.aliveWorkers++
			worker := newWorker()
			worker.Start(wp.workerQueue, wp.onPanic)
			wp.workerQueue <- worker
		}
	}
}

// Accept 接受任务
func (wp *WorkerPool) Accept(job Job) error {
	if job == nil {
		return errors.New("job cannot be nil")
	}
	wp.mux.Lock()
	select {
	case worker := <-wp.workerQueue:
		// 有闲置的worker
		wp.mux.Unlock()
		if worker != nil {
			worker.jobChannel <- job
		} else {
			return errors.New("workerPool is closed")
		}
	default:
		var worker *Worker
		if wp.aliveWorkers < wp.workerNum {
			// 还未达到Worker的上限
			wp.aliveWorkers++
			wp.mux.Unlock()
			worker = newWorker()
			worker.Start(wp.workerQueue, wp.onPanic)
		} else if wp.aliveWorkers == wp.workerNum {
			wp.mux.Unlock()
			// 这里可能会阻塞，会阻塞的情况为：当前所有的worker都还没执行完任务就又来了新的协程任务
			fmt.Println("worker number equal to alive number")
			worker = <-wp.workerQueue
			fmt.Println("now get worker from workerQueue")
		} else if wp.aliveWorkers > wp.workerNum {
			wp.mux.Unlock()
			return errors.New("worker number less than alive number")
		}
		// wp.AliveNum == wp.WorkerNum 可能会出现worker = nil
		if worker != nil {
			worker.jobChannel <- job
		} else {
			return errors.New("worker pool has been closed")
		}
	}
	return nil
}

func (wp *WorkerPool) Len() uint16 {
	wp.mux.RLock()
	defer wp.mux.RUnlock()
	return wp.aliveWorkers
}

func (wp *WorkerPool) Close() {
	wp.closeOnce.Do(func() {
		wp.mux.Lock()
		wp.closed = true
		wp.workerNum = 0
		wp.maxWorkers = 0
		for 0 < wp.aliveWorkers {
			worker := <-wp.workerQueue
			worker.Close()
			wp.aliveWorkers--
		}
		close(wp.workerQueue)
		wp.mux.Unlock()
	})
}
