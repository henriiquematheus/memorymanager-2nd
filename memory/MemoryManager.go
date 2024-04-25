package memory

import (
	"fmt"
	"math"
	"memory-manager-simulator/common"
	"memory-manager-simulator/process"
)

type MemoryManager struct {
	physicalMemory [][]*common.SubProcess
	logicalMemory  map[string][]*FrameMemory
	pageSize       int
}

func NewMemoryManager(pageSize int, physicalMemorySize int) *MemoryManager {
	pages := physicalMemorySize / pageSize
	return &MemoryManager{
		physicalMemory: make([][]*common.SubProcess, pages),
		logicalMemory:  make(map[string][]*FrameMemory),
		pageSize:       pageSize,
	}
}

func DefaultMemoryManager() *MemoryManager {
	return NewMemoryManager(4, 256)
}

func (mm *MemoryManager) Write(p *process.Process) {
	mm.writeWithPaging(p)
}

func (mm *MemoryManager) writeWithPaging(p *process.Process) {
	subProcesses := p.GetSubProcesses()
	frames := mm.findFramePages(len(subProcesses))
	spaces := int(math.Ceil(float64(len(p.GetSubProcesses())) / float64(mm.pageSize)))
	subProcesses := p.GetSubProcesses()

	if spaces <= len(frames) {
		spIndex := 0
		for i := 0; i < spaces; i++ {
			fm := frames[i]
			for j := 0; j < mm.pageSize; j++ {
				if spIndex < len(subProcesses) {

					sp := subProcesses[spIndex]
					mm.physicalMemory[fm.PageNumber][j] = sp
					fm.SetDisplacement(j)

					mm.logicalMemory[sp.GetID()] = append(mm.logicalMemory[sp.GetID()], fm)
					spIndex++
				} else {
					break
				}
			}
		}
	} else {
		fmt.Println("Não há espaço suficiente na memória")
	}
}

func (mm *MemoryManager) findFramePages(subProcessCount int) []*FrameMemory {
	var frames []*FrameMemory
	for i := 0; i < len(mm.physicalMemory); i++ {

		if mm.physicalMemory[i][0] == nil {
			frames = append(frames, NewFrameMemory(i, 0))
		}
	}
	return frames
}
