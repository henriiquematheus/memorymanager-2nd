package memory

// AddressMemory representa um intervalo de endereços de memória.
type AddressMemory struct {
	Start int // Início do intervalo
	End   int // Fim do intervalo
}

// GetStart retorna o início do intervalo.
func (am *AddressMemory) GetStart() int {
	return am.Start
}

// SetStart define o início do intervalo.
func (am *AddressMemory) SetStart(start int) {
	am.Start = start
}

// GetEnd retorna o fim do intervalo.
func (am *AddressMemory) GetEnd() int {
	return am.End
}

// SetEnd define o fim do intervalo.
func (am *AddressMemory) SetEnd(end int) {
	am.End = end
}

// GetSize retorna o tamanho do intervalo.
func (am *AddressMemory) GetSize() int {
	return (am.End - am.Start) + 1
}
