package interfaces

type SubProcess interface {
	GetID() string
	GetProcessID() string
	GetInstructionsExecuted() int
	SetInstructionsExecuted(instructionsExecuted int)
	GetInstructionsNumber() int
}

type SystemOperationHandler interface {
	SystemCall(callType SystemCallType, arg interface{}) interface{}
}
type ProcessListener interface {
	CoreExecuted(coreID int, processID string)
	ClockExecuted(clockTime int)
}

// CoreManager define a interface para gerenciar núcleos de CPU.
type CoreManager interface {
	Execute(process SubProcess)
	FinishExecution(process SubProcess)
	IsFinished() bool
	ExecuteProcesses()
	GetCores() []Core
	RegisterProcess(coreID int, process SubProcess) // Changed this line
}

// AddressMemory é uma interface para manipulação de endereços de memória
type AddressMemory interface {
	GetStart() int
	SetStart(start int)
	GetEnd() int
	SetEnd(end int)
	GetSize() int
}

type Core interface {
	Run()
	IsEmpty() bool
	GetActuallyProcess() SubProcess
	SetActuallyProcess(actuallyProcess SubProcess)
	GetID() int
	SetID(id int)
}
