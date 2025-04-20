package gopool

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

type Score struct {
	num int
}

func (s *Score) Do() error {
	if s.num%10000 == 0 {
		log.Default().Println("num", s.num)
	}
	if s.num%2 == 0 {
		panic(s.num)
	}
	time.Sleep(time.Second * 5)
	return nil
}

// go test -v -test.run TestWorkerPool_Run
func TestWorkerPool_Run(t *testing.T) {
	log.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
	p := NewWorkerPool(1000, 1100)
	p.OnPanic(func(msg interface{}) {
		log.Println("error:", msg)
	})
	datanum := 100 * 100 * 100
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		start := time.Now()
		for i := 1; i <= datanum; i++ {
			sc := &Score{num: i}
			if err := p.Accept(sc); err != nil {
				fmt.Println("err:\t", err)
				break
			}
			if i%10000 == 0 {
				log.Println("send num:", i)
			}
			//randNum := rand.Intn(10) + 1000
			//log.Println(randNum)
			//p.AdjustSize(uint16(randNum))
		}
		log.Println("start wait.....")
		p.Close()
		log.Printf("cost time:%s\n", time.Since(start).String())
		log.Println("stop over.....")
		log.Println("the last runtime.NumGoroutine() :", runtime.NumGoroutine())
	}()
	for {
		time.Sleep(1 * time.Second)
		log.Printf("runtime.NumGoroutine() :%d\n", runtime.NumGoroutine())
	}
}

func TestWorkerPool_Pause(t *testing.T) {
	log.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
	p := NewWorkerPool(5, 1100)
	p.OnPanic(func(msg interface{}) {
		log.Println("error:", msg)
	})
	datanum := 100
	var poolNum uint16
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		start := time.Now()
		go func() {
			for p.Len() > 0 {
			}
			time.Sleep(2 * time.Second)
			p.Resume(poolNum)
		}()
		for i := 1; i <= datanum; i++ {
			sc := &Score{num: i}
			if err := p.Accept(sc); err != nil {
				fmt.Println("err:\t", err)
				break
			}
			if i == 2 {
				log.Println("send num:", i)
				poolNum = p.Len()
				p.Pause()
			}
		}
		log.Println("start wait.....")
		p.Close()
		log.Printf("cost time:%s\n", time.Since(start).String())
		log.Println("stop over.....")
		log.Println("the last runtime.NumGoroutine() :", runtime.NumGoroutine())
	}()

	for {
		time.Sleep(1 * time.Second)
		log.Printf("runtime.NumGoroutine() :%d\n", runtime.NumGoroutine())
	}
}
