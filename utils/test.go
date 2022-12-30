package utils

import (
	"fmt"
	"testing"
	"time"
)

func ShowError(t *testing.T, expected interface{}, got interface{}) {
	expectedMessage := fmt.Sprintf("%vExpected: %v%v", Colours.Green, expected, Colours.Reset)
	gotMessage := fmt.Sprintf("%vGot: %v%v", Colours.Red, got, Colours.Reset)

	t.Errorf("\nError \n%v \n%v", expectedMessage, gotMessage)
}

func ErrorWhenNoMessageInChannelInGivenTime[T any](t *testing.T, ch chan T, duration time.Duration) {
	timeoutMessage := fmt.Sprintf("%vTimeout", Colours.Yellow)

	select {
	case <-ch:
		return
	case <-time.After(duration):
		ShowError(t, true, timeoutMessage)
	}
}

func ConstructRunWithCleanup(cleanUp func()) func(T *testing.T, description string, testFunction func(t *testing.T)) {
	return func(T *testing.T, description string, testFunction func(t *testing.T)) {
		T.Run(description, func(t *testing.T) {
			testFunction(t)
			cleanUp()
		})
	}
}
