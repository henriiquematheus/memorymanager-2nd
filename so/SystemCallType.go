package so

type SystemCallType int

const (
	OPEN_PROCESS SystemCallType = iota
	READ_PROCESS
	CLOSE_PROCESS
	CREATE_PROCESS
	WRITE_PROCESS
)

func GetSystemCallTypeName(callType SystemCallType) string {
	switch callType {
	case OPEN_PROCESS:
		return "Open Process"
	case READ_PROCESS:
		return "Read Process"
	case CLOSE_PROCESS:
		return "Close Process"
	case CREATE_PROCESS:
		return "Create Process"
	case WRITE_PROCESS:
		return "Write Process"
	default:
		return "Unknown System Call Type"
	}
}

//
