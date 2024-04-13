package scheduler

import (
	"memory-manager-simulator/cpu"
	"memory-manager-simulator/process"
)

type Scheduler struct {
	cpu *cpu.CpuManager
}

func NewScheduler(processListener process.ProcessListener) *Scheduler {
	return &Scheduler{
		cpu: cpu.NewCpuManager(processListener),
	}
}

func NewDefaultScheduler() *Scheduler {
	return &Scheduler{
		cpu: cpu.NewCpuManager(nil),
	}
}

func (s *Scheduler) Execute(p *process.SOProcess) {

	if !p.IsStarted() {

		p.Start()
	}

	if p.IsReady() {

		s.cpu.Execute(p)
	} else {

		p.AdjustPriority()
	}
}

func (s *Scheduler) CloseProcess(p *process.SOProcess) {

	s.cpu.FinishExecution(p)

	p.Finish()
}

func (s *Scheduler) IsFinished() bool {

	return s.cpu.IsFinished()
}

//
