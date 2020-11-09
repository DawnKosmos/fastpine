package helper

import "time"

func DateToTime(day, month, year string) int64 {
	t, _ := time.Parse("01 02 2006 15:04 MST", month+" "+day+" "+year+" 00:00 GMT")
	return t.Unix()
}

func DateToTimeHourly(day, month, year, hour string) int64 {
	t, _ := time.Parse("01 02 2006 15:04 MST", month+" "+day+" "+year+" "+hour+" GMT")
	return t.Unix()
}
