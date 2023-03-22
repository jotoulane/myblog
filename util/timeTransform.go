package util

import (
	"time"
)

func TimeTransform(myTime time.Time) string {
	newLayout := "2006-01-02 15:04:05"
	return myTime.Format(newLayout)

}
