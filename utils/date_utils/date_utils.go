package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow returns a UTC date of the current time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns a UTC string representation of the current time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDbFormat
func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
