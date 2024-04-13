package cpu

import (
	"fmt" // Importe o novo pacote de interfaces
	"memory-manager-simulator/interfaces"
)

type Core struct {
	id                    int
	instructionsPerSecond int
	actuallyProcess       interfaces.SubProcess
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

func (c *Core) IsEmpty() bool {
	return c.actuallyProcess == nil
}

func (c *Core) GetActuallyProcess() interfaces.SubProcess {
	return c.actuallyProcess
}

func (c *Core) SetActuallyProcess(actuallyProcess interfaces.SubProcess) { // Esta linha já foi alterada
	c.actuallyProcess = actuallyProcess
}

func (c *Core) finishExecution() {
	c.actuallyProcess = nil
}

func (c *Core) GetID() int {
	return c.id
}

func (c *Core) SetID(id int) {
	c.id = id
}
