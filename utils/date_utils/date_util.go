package date_utils

import "time"

const (
	apiDateLayout = "Monday, 02 January 2006 15:04:05 MST"
	apiDBLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	loc,_ := time.LoadLocation("Africa/Dar_es_Salaam")
	return  time.Now().In(loc)
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}