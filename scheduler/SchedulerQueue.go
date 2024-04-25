package scheduler

import (
	"container/heap"
	"memory-manager-simulator/cpu"
	"memory-manager-simulator/process"
	"sync"
)

type SchedulerQueue struct {
	*Scheduler
	queue        PriorityQueue
	subProcesses map[string][]*process.SubProcess
	mutex        sync.Mutex
}

func NewSchedulerQueue(comp Comparator) *SchedulerQueue {
	sq := &SchedulerQueue{
		Scheduler:    NewDefaultScheduler(),
		queue:        make(PriorityQueue, 0),
		subProcesses: make(map[string][]*process.SubProcess),
	}
	heap.Init(&sq.queue, comp)
	sq.cpu = cpu.NewCpuManager(sq)
	return sq
}

func (sq *SchedulerQueue) Execute(p *process.SOProcess) {
	sps := SystemOperation.SystemCall(SystemCallTypeReadProcess, p)
	heap.Push(&sq.queue, p)
	sq.subProcesses[p.GetID()] = sps
	sq.registerInProcessor(p.GetID())
}

func (sq *SchedulerQueue) CloseProcess(p *process.SOProcess) {
	delete(sq.subProcesses, p.GetID())
}

func (sq *SchedulerQueue) IsFinished() bool {
	return sq.queue.Len() == 0
}

func (sq *SchedulerQueue) registerInProcessor(processID string) {
	for _, core := range sq.cpu.GetCores() {
		if core.GetActuallyProcess() == nil {
			sq.coreExecuted(core.GetID(), processID)
			break
		}
	}
}

func (sq *SchedulerQueue) coreExecuted(coreID int, processID string) {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()

	if sps, ok := sq.subProcesses[processID]; ok {
		if len(sps) == 0 {
			sq.queue.Pop()
			delete(sq.subProcesses, processID)
		}
		actuallyProcess := sq.queue.Peek().(*process.SOProcess)
		if len(sps) > 0 {
			actuallySubProcess := sps[len(sps)-1]
			sps = sps[:len(sps)-1]
			sq.cpu.RegisterProcess(coreID, actuallySubProcess)
		}
	}
}

type PriorityQueue []*process.SOProcess

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority() < pq[j].Priority() }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*process.SOProcess)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
