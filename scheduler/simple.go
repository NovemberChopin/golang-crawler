package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workerChan <- request
	}()
}

// 把初始请求发送给 Scheduler
func (s *SimpleScheduler) ConfigMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}
