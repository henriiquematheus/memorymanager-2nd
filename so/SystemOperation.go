package so

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/process"
	"memory-manager-simulator/scheduler"
)

type SystemOperation struct {
	memoryManager *memory.MemoryManager
	scheduler     *scheduler.Scheduler
}

func NewSystemOperation() *SystemOperation {
	return &SystemOperation{
		memoryManager: memory.NewMemoryManager(0, 0),
		scheduler:     scheduler.NewDefaultScheduler(),
	}
}

// Função SystemCall para lidar com chamadas de sistema
func (so *SystemOperation) SystemCall(callType SystemCallType, arg interface{}) interface{} {
	switch callType {
	case CREATE_PROCESS:
		if so.memoryManager == nil {
			so.memoryManager = memory.NewMemoryManager(0, 0)
		}
		if so.scheduler == nil {
			so.scheduler = scheduler.NewDefaultScheduler()
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
	case CLOSE_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}

		so.memoryManager.Close(p)
		so.scheduler.CloseProcess(p)
		return nil
	case READ_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}

		return so.memoryManager.Read(p)
	case OPEN_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}
		so.scheduler.Execute(p)
		return nil
	default:
		return nil
	}
	return nil
}
