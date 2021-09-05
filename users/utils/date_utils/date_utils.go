package dateutils

import "time"

const (
	apiDatelayout = "2006-01-02T15:05:05Z"
)

func GetNowString() string {
	return GetNow().Format(apiDatelayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
