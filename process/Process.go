package process

import (
	"fmt"
	"memory-manager-simulator/interfaces" // Importando o pacote interfaces
)

type Process struct {
	ID                   string
	SizeInMemory         int
	AddressMemory        interfaces.AddressMemory // Agora usa a interface
	CurrentFragment      int
	SubProcesses         []interfaces.SubProcess // Agora usa a interface
	Started              bool                    // Novo campo para controlar se o processo foi iniciado
	Ready                bool                    // Novo campo para controlar se o processo está pronto
	InstructionsExecuted int
	InstructionsNumber   int
	Priority             int
}

var (
	processCounter = 0
)

func NewProcess(sizeInMemory int) *Process {
	processCounter++

	p := &Process{
		ID:              fmt.Sprintf("P%d", processCounter),
		SizeInMemory:    sizeInMemory,
		CurrentFragment: 1,
		SubProcesses:    make([]interfaces.SubProcess, 0), // Inicializando slice de subprocessos
	}

	return p
}

func (p *Process) IncrementFragment() {
	p.CurrentFragment++
}

func (p *Process) GetID() string {
	return p.ID
}

func (p *Process) SetID(id string) {
	p.ID = id
}

func (p *Process) GetSizeInMemory() int {
	return p.SizeInMemory
}

func (p *Process) SetSizeInMemory(sizeInMemory int) {
	p.SizeInMemory = sizeInMemory
}

func (p *Process) GetAddressMemory() interfaces.AddressMemory {
	return p.AddressMemory
}

func (p *Process) SetAddressMemory(am interfaces.AddressMemory) {
	p.AddressMemory = am
}

// Implementação do método IsStarted para verificar se o processo foi iniciado
func (p *Process) IsStarted() bool {
	return p.Started
}

// Implementação do método Start para iniciar o processo
func (p *Process) Start() {
	p.Started = true
}
func (p *Process) GetInstructionsExecuted() int {
	return p.InstructionsExecuted
}

func (p *Process) SetInstructionsExecuted(instructionsExecuted int) {
	p.InstructionsExecuted = instructionsExecuted
}

func (p *Process) GetInstructionsNumber() int {
	return p.InstructionsNumber
}

// Implementação do método IsReady para verificar se o processo está pronto
func (p *Process) IsReady() bool {
	return p.Ready
}

// Implementação do método AdjustPriority para ajustar a prioridade do processo
func (p *Process) AdjustPriority() {
	// Implementação para ajustar a prioridade do processo
}

// Implementação do método Finish para finalizar o processo
func (p *Process) Finish() {
	// Implementação para finalizar o processo
	p.Ready = false
}

// Método para obter os subprocessos do processo
func (p *Process) GetSubProcesses() []interfaces.SubProcess {
	return p.SubProcesses
}

// Método para adicionar um subprocesso ao processo
func (p *Process) AddSubProcess(subProcess interfaces.SubProcess) {
	p.SubProcesses = append(p.SubProcesses, subProcess)
}

func (p *Process) GetProcessID() string {
	return p.ID
}

func (p *Process) GetPriority() int {
	return p.Priority
}
