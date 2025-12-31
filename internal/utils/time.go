package utils

import "time"

// Gets the current system time and returns it as a usable string.
func GetTime() string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return timestamp
}
