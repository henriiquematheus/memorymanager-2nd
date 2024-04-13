package memory

type AddressMemory struct {
	Start int // Início do intervalo
	End   int // Fim do intervalo
}

func (am *AddressMemory) GetStart() int {
	return am.Start
}

func (am *AddressMemory) SetStart(start int) {
	am.Start = start
}

func (am *AddressMemory) GetEnd() int {
	return am.End
}

func (am *AddressMemory) SetEnd(end int) {
	am.End = end
}

func (am *AddressMemory) GetSize() int {
	return (am.End - am.Start) + 1
}
