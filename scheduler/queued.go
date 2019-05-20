package scheduler

import (
	"crawler/engine"
	"fmt"
)

// 使用队列来调度任务

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	// 对于队列实现来讲，每个 worker 共用一个 channel
	return make(chan engine.Request)
}

// 提交请求任务到 requestChan
func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

// 告诉外界有一个 worker 可以接收 request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		// 创建请求队列和工作队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan: // 当 requestChan 收到数据
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 当 workerChan 收到数据
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 当请求队列和认读队列都不为空时，给任务队列分配任务
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
				fmt.Printf("requestQ: %d, workerQ: %d\n", len(requestQ), len(workerQ))
			}
		}
	}()
}
