package scheduler

import (
	"container/heap"
	"memory-manager-simulator/cpu"
	"memory-manager-simulator/process"
	"memory-manager-simulator/so"
	"sync"
)

type Comparator func(p1, p2 *process.Process) bool

type SchedulerQueue struct {
	*Scheduler
	queue        PriorityQueue
	subProcesses map[string][]*process.SubProcess
	mutex        sync.Mutex
	systemOp     *so.SystemOperation // add a field for SystemOperation
}

func NewSchedulerQueue(comp Comparator) *SchedulerQueue {
	sq := &SchedulerQueue{
		Scheduler:    NewDefaultScheduler(),
		queue:        make(PriorityQueue, 0),
		subProcesses: make(map[string][]*process.SubProcess),
		systemOp:     so.NewSystemOperation(), // initialize SystemOperation
	}
	heap.Init(&sq.queue)
	sq.cpu = cpu.NewCpuManager(sq)
	return sq
}

func (sq *SchedulerQueue) Execute(p *process.Process) {
	sps := sq.systemOp.SystemCall(so.READ_PROCESS, p) // use the instance to call the method
	heap.Push(&sq.queue, p)
	sq.subProcesses[p.GetID()] = sps.([]*process.SubProcess) // make sure to cast the result to the correct type
	sq.RegisterInProcessor(p.GetID())
}

func (sq *SchedulerQueue) RegisterInProcessor(processID string) {
	for _, core := range sq.cpu.GetCores() {
		if core.GetActuallyProcess() == nil {
			sq.CoreExecuted(core.GetID(), processID)
			break
		}
	}
}

func (sq *SchedulerQueue) CoreExecuted(coreID int, processID string) {
	sq.mutex.Lock()
	defer sq.mutex.Unlock()

	if sps, ok := sq.subProcesses[processID]; ok {
		if len(sps) == 0 {
			sq.queue.Pop()
			delete(sq.subProcesses, processID)
		}
		actuallyProcess := sq.queue.Peek().(*process.Process)
		if len(sps) > 0 {
			actuallySubProcess := sps[len(sps)-1]
			sps = sps[:len(sps)-1]
			sq.cpu.RegisterProcess(coreID, actuallySubProcess)
		}
	}
}

func (sq *SchedulerQueue) ClockExecuted(clockTime int) {
	// Implemente a lógica aqui
}

type PriorityQueue []*process.Process

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority() > pq[j].Priority() }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*process.Process)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() interface{} {
	return (*pq)[0]
}
