package cpu

import (
	"errors"
	"fmt"
	"memory-manager-simulator/common"
	"memory-manager-simulator/scheduler"
	"time"
)

type CpuManager struct {
	cores           []*Core
	clockInterval   time.Duration
	processListener scheduler.ProcessListener
}

func NewCpuManager(processListener scheduler.ProcessListener) *CpuManager {
	return &CpuManager{
		cores:           make([]*Core, 0),
		clockInterval:   0,
		processListener: processListener,
	}
}

func (cm *CpuManager) RegisterProcess(coreIndex int, sp *common.SubProcess) error {
	if coreIndex < 0 || coreIndex >= len(cm.cores) {
		return errors.New("índice de núcleo inválido")
	}
	cm.cores[coreIndex].SetActuallyProcess(sp)
	return nil
}

func (cm *CpuManager) Clock() {

	ticker := time.NewTicker(cm.clockInterval)
	go func() {
		for range ticker.C {
			fmt.Println("****")
			cm.ExecuteProcesses()
		}
	}()
}

func (cm *CpuManager) ExecuteProcesses() {
	for _, core := range cm.cores {

		if !core.IsEmpty() {

			core.Run()

			cm.processListener.CoreExecuted(core.GetID(), core.GetActuallyProcess().GetID())
		}
	}
}
