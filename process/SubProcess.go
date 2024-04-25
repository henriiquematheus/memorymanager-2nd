package process

import (
	"fmt"
)

type SubProcess struct {
	Instructions         int
	ID                   string
	ProcessID            string
	InstructionsExecuted int
}

var subProcessCount int

func NewSubProcess(processID string, instructions int) *SubProcess {
	subProcessCount++
	id := fmt.Sprintf("%s%d", processID, subProcessCount)
	return &SubProcess{
		Instructions: instructions,
		ID:           id,
		ProcessID:    processID,
	}
}

func (sp *SubProcess) GetInstructions() int {
	return sp.Instructions
}

func (sp *SubProcess) SetInstructions(instructions int) {
	sp.Instructions = instructions
}

func (sp *SubProcess) GetID() string {
	return sp.ID
}

func (sp *SubProcess) SetID(id string) {
	sp.ID = id
}

func (sp *SubProcess) GetProcessID() string {
	return sp.ProcessID
}

func (sp *SubProcess) SetProcessID(processID string) {
	sp.ProcessID = processID
}

func (sp *SubProcess) GetInstructionsExecuted() int {
	return sp.InstructionsExecuted
}

func (sp *SubProcess) SetInstructionsExecuted(executed int) {
	sp.InstructionsExecuted = executed
}

func (sp *SubProcess) GetInstructionsNumber() int {
	return sp.Instructions
}
