package scheduler

import "memory-manager-simulator/common"

type SJFScheduler struct {
	processQueue []*common.Process
}

func NewSJFScheduler() *SJFScheduler {
	return &SJFScheduler{
		processQueue: make([]*common.Process, 0),
	}
}

func (s *SJFScheduler) AddProcess(p *common.Process) {
	s.processQueue = append(s.processQueue, p)
}

func (s *SJFScheduler) Schedule() *common.Process {
	if len(s.processQueue) == 0 {
		return nil
	}
	shortestProcess := s.processQueue[0]
	for _, p := range s.processQueue {
		if p.ExecutionTime < shortestProcess.ExecutionTime {
			shortestProcess = p
		}
	}
	return shortestProcess
}
