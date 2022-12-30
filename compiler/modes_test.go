package compiler

import (
	"go-checker/utils"
	"testing"
)

type UseCaseMock struct {
}

var compileProgramCallsCount = 0

func (_ UseCaseMock) CompileProgram(programName string, sendDone chan bool) {
	compileProgramCallsCount++
	sendDone <- true
}

type PresenterMock struct {
}

var hasPresenterReceivedMessage = make(chan bool)

func (_ PresenterMock) PresentCompilation(done chan bool) {
	hasPresenterReceivedMessage <- <-done
}
func (_ PresenterMock) PresentCompilationSuccess() {

}

func getAppToAppModeInstance() AppToAppMode {
	return ConstructAppToAppMode(UseCaseMock{}, PresenterMock{})
}

func cleanUpTestAppToAppMode() {
	close(hasPresenterReceivedMessage)
	hasPresenterReceivedMessage = make(chan bool)
	compileProgramCallsCount = 0
}

func TestAppToAppMode(T *testing.T) {

	run := utils.ConstructRunWithCleanup(cleanUpTestAppToAppMode)

	run(T, "Should stop presenting compilation When compilation is ended", func(t *testing.T) {
		testMode := getAppToAppModeInstance()
		testMode.Compile()

		receivedMessage := <-hasPresenterReceivedMessage

		if !receivedMessage {
			utils.ShowError(t, true, false)
		}
		hasAllThreeCompilationsFinished := compileProgramCallsCount == 3
		if !hasAllThreeCompilationsFinished {
			utils.ShowError(t, 3, compileProgramCallsCount)
		}
	})
}
