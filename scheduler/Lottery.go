package scheduler

import (
	"math/rand"
	"memory-manager-simulator/common"
)

// LotteryScheduler é uma estrutura que implementa o escalonamento por loteria.
type LotteryScheduler struct {
	processes []*common.Process // Lista de processos disponíveis para sorteio
}

// NewLotteryScheduler cria uma nova instância de LotteryScheduler.
func NewLotteryScheduler() *LotteryScheduler {
	return &LotteryScheduler{
		processes: make([]*common.Process, 0),
	}
}

// AddProcess adiciona um novo processo à lista de processos disponíveis para sorteio.
func (l *LotteryScheduler) AddProcess(p *common.Process) {
	l.processes = append(l.processes, p)
}

// Schedule seleciona aleatoriamente um processo da lista de processos disponíveis.
func (l *LotteryScheduler) Schedule() *common.Process {
	if len(l.processes) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(l.processes))
	nextProcess := l.processes[randomIndex]
	return nextProcess
}
