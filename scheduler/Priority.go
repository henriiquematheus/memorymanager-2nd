package scheduler

import "memory-manager-simulator/common"

// PriorityScheduler é uma estrutura que implementa o escalonamento por prioridade.
type PriorityScheduler struct {
	processQueue []*common.Process // Fila de processos a serem escalonados
}

// NewPriorityScheduler cria uma nova instância de PriorityScheduler.
func NewPriorityScheduler() *PriorityScheduler {
	return &PriorityScheduler{
		processQueue: make([]*common.Process, 0),
	}
}

// AddProcess adiciona um novo processo à fila.
func (p *PriorityScheduler) AddProcess(process *common.Process) {
	p.processQueue = append(p.processQueue, process)
}

// Schedule seleciona o próximo processo na fila com a maior prioridade.
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
