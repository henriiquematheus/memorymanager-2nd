package scheduler

import (
	"memory-manager-simulator/cpu"
	"memory-manager-simulator/interfaces"
	"memory-manager-simulator/process"
)

type Scheduler struct {
	cpu interfaces.CoreManager
}

func NewScheduler(processListener interfaces.ProcessListener) *Scheduler {
	return &Scheduler{
		cpu: cpu.NewCpuManager(processListener),
	}
}

func NewDefaultScheduler() *Scheduler {
	return &Scheduler{
		cpu: cpu.NewCpuManager(nil),
	}
}

func (s *Scheduler) Execute(p *process.Process) {
	if !p.IsStarted() {
		p.Start()
	}
	if p.IsReady() {
		s.cpu.Execute(p)
	} else {
		p.AdjustPriority()
	}
}

func (s *Scheduler) CloseProcess(p *process.Process) {
	s.cpu.FinishExecution(p)
	p.Finish()
}

func (s *Scheduler) IsFinished() bool {
	return s.cpu.IsFinished()
}
