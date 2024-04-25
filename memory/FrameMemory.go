package memory

type FrameMemory struct {
	PageNumber   int
	Displacement int
}

func NewFrameMemory(pageNumber, displacement int) *FrameMemory {
	return &FrameMemory{
		PageNumber:   pageNumber,
		Displacement: displacement,
	}
}

func (fm *FrameMemory) GetPageNumber() int {
	return fm.PageNumber
}

func (fm *FrameMemory) SetPageNumber(pageNumber int) {
	fm.PageNumber = pageNumber
}

func (fm *FrameMemory) GetDisplacement() int {
	return fm.Displacement
}

func (fm *FrameMemory) SetDisplacement(displacement int) {
	fm.Displacement = displacement
}
