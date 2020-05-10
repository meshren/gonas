package utils

import (
	"time"
)

var location *time.Location

func init() {
	location = time.FixedZone("GMT", 8*3600)       // 东八
}

func AddTime(t1 time.Time,duration string) (t time.Time, err error) {
	t2, err := time.ParseDuration(duration)
	if err != nil {
		return
	}
	t = t1.Add(t2)
	return
}
