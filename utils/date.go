package utils

import (
	"time"
)

// GetCurrTime get current timestamp
func GetCurrTime() int {
	return int(time.Now().Unix())
}