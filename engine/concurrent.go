package engine

import (
	"log"
)

// 并发引擎
type ConcurrendEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// 任务调度器
type Scheduler interface {
	Submit(request Request) // 提交任务
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(w chan Request)
	Run()
}

func (e *ConcurrendEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建 goruntine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	// engine把请求任务提交给 Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		// 接受 Worker 的解析结果
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v\n", itemCount, item)
			itemCount++
		}

		// 然后把 Worker 解析出的 Request 送给 Scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in) // 告诉调度器任务空闲
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
