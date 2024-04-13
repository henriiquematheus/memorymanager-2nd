package scheduler

type ProcessListener interface {
	CoreExecuted(coreID int, processID string)
	ClockExecuted(clockTime int)
}

//
