package process

import (
	"fmt"
)

type Process struct {
	ID              string
	SizeInMemory    int
	AddressMemory   AddressMemory
	CurrentFragment int
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

func (p *Process) GetAddressMemory() AddressMemory {
	return p.AddressMemory
}

func (p *Process) SetAddressMemory(am AddressMemory) {
	p.AddressMemory = am
}
