package etc

import "time"

func GetCurrentTime() string {
	// This function returns the current time in a specific format.
	// The format is "2006-01-02 15:04:05" which is the standard time format in Go.
	return time.Now().Format("2006-01-02 15:04:05")
}
