package scheduler

import "memory-manager-simulator/common"

// FCFSScheduler é uma estrutura que implementa o escalonamento FCFS.
type FCFSScheduler struct {
	processQueue []*common.Process // Fila de processos a serem escalonados
}

// NewFCFSScheduler cria uma nova instância de FCFSScheduler.
func NewFCFSScheduler() *FCFSScheduler {
	return &FCFSScheduler{
		processQueue: make([]*common.Process, 0),
	}
}

// AddProcess adiciona um novo processo à fila.
func (f *FCFSScheduler) AddProcess(p *common.Process) {
	f.processQueue = append(f.processQueue, p)
}

// Schedule seleciona o próximo processo na fila de acordo com a política FCFS.
func (f *FCFSScheduler) Schedule() *common.Process {
	if len(f.processQueue) == 0 {
		return nil
	}
	nextProcess := f.processQueue[0]
	f.processQueue = f.processQueue[1:]
	return nextProcess
}
