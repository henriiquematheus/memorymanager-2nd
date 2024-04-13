package scheduler

import "memory-manager-simulator/common"

type PriorityScheduler struct {
	processQueue []*common.Process
}

func NewPriorityScheduler() *PriorityScheduler {
	return &PriorityScheduler{
		processQueue: make([]*common.Process, 0),
	}
}

func (p *PriorityScheduler) AddProcess(process *common.Process) {
	p.processQueue = append(p.processQueue, process)
}

func (p *PriorityScheduler) Schedule() *common.Process {
	if len(p.processQueue) == 0 {
		return nil
	}
	highestPriorityProcess := p.processQueue[0]
	for _, process := range p.processQueue {
		if process.Priority > highestPriorityProcess.Priority {
			highestPriorityProcess = process
		}
	}
	return highestPriorityProcess
}

//
