package custom_date

import "time"

func Get2DaysFromNow() string {
	return time.Now().Add(time.Hour * 48).Format("2006-01-02 15:04:05")
}
