package so

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/process"
)

type SystemOperation struct {
	memoryManager *memory.MemoryManager
	scheduler     *Scheduler
}

func NewSystemOperation() *SystemOperation {
	return &SystemOperation{
		memoryManager: memory.NewMemoryManager(0),
		scheduler:     NewScheduler(),
	}
}

func (so *SystemOperation) SystemCall(callType SystemCallType, arg interface{}) interface{} {
	switch callType {
	case CREATE_PROCESS:
		if so.memoryManager == nil {
			so.memoryManager = memory.NewMemoryManager(0)
		}
		if so.scheduler == nil {
			so.scheduler = NewScheduler()
		}
		processSize, ok := arg.(int)
		if !ok {
			return nil
		}
		return process.NewProcess(processSize)
	case WRITE_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}
		so.memoryManager.Write(p)
	}
	return nil
}
