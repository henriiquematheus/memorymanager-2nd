package cpu

import (
	"fmt"
	"memory-manager-simulator/process"
)

// Core representa um núcleo de CPU.
type Core struct {
	id                    int
	instructionsPerSecond int
	actuallyProcess       *process.SubProcess
}

func NewCore(instructionsPerSecond int) *Core {
	return &Core{
		instructionsPerSecond: instructionsPerSecond,
	}
}

func (c *Core) Run() {
	if c.actuallyProcess == nil {
		return
	}

	sumInstructionsExecuted := c.actuallyProcess.GetInstructionsExecuted() + c.instructionsPerSecond
	c.actuallyProcess.SetInstructionsExecuted(sumInstructionsExecuted)

	fmt.Printf("Núcleo %d Executando Subprocesso %s\n", c.id, c.actuallyProcess.GetID())
	if c.actuallyProcess.GetInstructionsExecuted() >= c.actuallyProcess.GetInstructionsNumber() {
		c.finishExecution()
	}
}

// IsEmpty verifica se o núcleo está vazio.
func (c *Core) IsEmpty() bool {
	return c.actuallyProcess == nil
}

// GetActuallyProcess retorna o processo atualmente executado pelo núcleo.
func (c *Core) GetActuallyProcess() *process.SubProcess {
	return c.actuallyProcess
}

// SetActuallyProcess define o processo atualmente executado pelo núcleo.
func (c *Core) SetActuallyProcess(actuallyProcess *process.SubProcess) {
	c.actuallyProcess = actuallyProcess
}

// finishExecution finaliza a execução do processo atual.
func (c *Core) finishExecution() {
	c.actuallyProcess = nil
}

// GetID retorna o ID do núcleo.
func (c *Core) GetID() int {
	return c.id
}

// SetID define o ID do núcleo.
func (c *Core) SetID(id int) {
	c.id = id
}
