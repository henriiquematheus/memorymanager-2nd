package common

type Process struct {
	ID            string
	Size          int
	Status        string
	Priority      int
	ExecutionTime int
}

type SubProcess struct {
	ID   string
	Size int
}
