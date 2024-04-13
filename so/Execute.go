package so

import (
	"memory-manager-simulator/process"
)

func Execute(strategy int) {
	so := NewSystemOperation(strategy)

	pl := so.SystemCall(CREATE_PROCESS, 10).(*process.Process)
	so.SystemCall(WRITE_PROCESS, pl)

	p2 := so.SystemCall(CREATE_PROCESS, 20).(*process.Process)
	so.SystemCall(WRITE_PROCESS, p2)

	p3 := so.SystemCall(CREATE_PROCESS, 15).(*process.Process)
	so.SystemCall(WRITE_PROCESS, p3)
	//
	// Feche os processos p2 e p3
	so.SystemCall(CLOSE_PROCESS, p2)
	so.SystemCall(CLOSE_PROCESS, p3)

	p4 := so.SystemCall(CREATE_PROCESS, 25).(*process.Process)
	so.SystemCall(WRITE_PROCESS, p4)

	p5 := so.SystemCall(CREATE_PROCESS, 30).(*process.Process)
	so.SystemCall(WRITE_PROCESS, p5)

	p6 := so.SystemCall(CREATE_PROCESS, 18).(*process.Process)
	so.SystemCall(WRITE_PROCESS, p6)
}
