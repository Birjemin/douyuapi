package utils

import (
	"time"
)

// get current timestamp
func GetCurrTime() int {
	return int(time.Now().Unix())
}