package langserver

type LogLevel int

const (
	ERROR LogLevel = iota + 1
	WARN
	INFO
	DEBUG
	TRACE
)

// Set log level from integer value
func SetLogLevel(loglevel int) LogLevel {
	switch loglevel {
	case 1:
		return ERROR
	case 2:
		return WARN
	case 3:
		return INFO
	case 4:
		return DEBUG
	case 5:
		return TRACE
	}
	return ERROR
}
