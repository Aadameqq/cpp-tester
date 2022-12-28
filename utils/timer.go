package utils

import (
	"fmt"
	"time"
)

func StartTimer() func() string {
	start := time.Now()
	return func() string {
		return fmt.Sprintf("%f", time.Since(start).Seconds())
	}
}
