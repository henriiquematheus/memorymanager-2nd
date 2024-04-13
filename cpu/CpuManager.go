package cpu

import (
	"errors"
	"fmt"
	"memory-manager-simulator/interfaces"
	"memory-manager-simulator/process" // Ensure the correct package is imported here
	"time"
)

type CpuManager struct {
	cores           []interfaces.Core
	clockInterval   time.Duration
	processListener interfaces.ProcessListener
}

func NewCpuManager(processListener interfaces.ProcessListener) *CpuManager {
	return &CpuManager{
		cores:           make([]interfaces.Core, 0),
		clockInterval:   0,
		processListener: processListener,
	}
}

func (cm *CpuManager) Execute(sp interfaces.SubProcess) {
	for _, core := range cm.cores {
		if core.IsEmpty() {
			core.SetActuallyProcess(sp)
			break
		}
	}
}

func (cm *CpuManager) ExecuteProcesses() {
	for _, core := range cm.cores {
		if !core.IsEmpty() {
			core.Run()
			cm.processListener.CoreExecuted(core.GetID(), core.GetActuallyProcess().GetID())
		}
	}
}

// Changed the second parameter type to a pointer to interfaces.SubProcess
func (cm *CpuManager) RegisterProcess(coreID int, process interfaces.SubProcess) error {
	if coreID < 0 || coreID >= len(cm.cores) {
		return errors.New("invalid core index")
	}

	// Check if the process is a pointer to process.SubProcess
	subProcessPtr, ok := process.(*process.SubProcess)
	if !ok {
		return errors.New("failed to convert SubProcess to *process.SubProcess")
	}

	cm.cores[coreID].SetActuallyProcess(subProcessPtr)
	return nil
}

func (cm *CpuManager) FinishExecution(sp interfaces.SubProcess) {
	p, ok := sp.(*process.Process)
	if !ok {
		return
	}

	for _, core := range cm.cores {
		if !core.IsEmpty() && core.GetActuallyProcess().GetID() == p.GetID() {
			core.SetActuallyProcess(nil)
			break
		}
	}
}

func (cm *CpuManager) IsFinished() bool {
	for _, core := range cm.cores {
		if !core.IsEmpty() {
			return false
		}
	}
	return true
}

func (cm *CpuManager) GetCores() []interfaces.Core {
	return cm.cores
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
