package utils

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

// LevelColor returns the ANSI color code associated with the given log level.
func LevelColor(level Level) string {
	switch level {
	case TRACE:
		return White
	case DEBUG:
		return Gray
	case INFO:
		return White
	case WARN:
		return Yellow
	case ERROR:
		return Red
	case FATAL:
		return Red
	default:
		return White
	}
}
