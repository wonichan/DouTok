package gopool

import "fmt"

// 任务接口
type Job interface {
	Do() error
}

// worker 用来执行协程任务
type Worker struct {
	jobChannel chan Job
	quit       chan struct{}
}

func newWorker() *Worker {
	// 无缓冲channel，一个worker只能执行一个任务
	return &Worker{
		jobChannel: make(chan Job),
		quit:       make(chan struct{}),
	}
}

func (w *Worker) Start(wq chan<- *Worker, onPanic func(data interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				w.Start(wq, onPanic)
				fmt.Println("worker restarting")
				wq <- w
				fmt.Println("worker restarted")
				if onPanic != nil {
					onPanic(err)
				}
			}
		}()
		for {
			select {
			case job := <-w.jobChannel:
				// 执行任务
				job.Do()
				// 执行完任务后，将worker放回worker队列
				wq <- w
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Worker) Close() {
	w.quit <- struct{}{}
	close(w.quit)
	close(w.jobChannel)
	w.jobChannel = nil
}

func (w *Worker) IsAvailable() bool {
	return len(w.jobChannel) == 0 && w.jobChannel != nil
}

func (w *Worker) StopRun() {
	w.quit <- struct{}{}
}
