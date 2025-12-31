package utils

// Level represents the severity of a log message.
type Level string

const (
	TRACE Level = "TRACE"	// Very fine-grained logging
	DEBUG Level = "DEBUG"	// Debugging information
	INFO  Level = "INFO"	// General information
	WARN  Level = "WARN"	// Potential issues
	ERROR Level = "ERROR"	// Errors affecting operation
	FATAL Level = "FATAL"	// Unrecoverable errors
)
