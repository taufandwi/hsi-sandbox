package custom_date

import "time"

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
