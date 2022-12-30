package compiler

import (
	"go-checker/utils"
	"testing"
)

type CompilationRunnerMock struct {
}

func (cr CompilationRunnerMock) Run(programName string) {

}

func getUseCaseInstance() UseCase {
	return ConstructUseCase(CompilationRunnerMock{})
}

func TestUseCase(testRunner *testing.T) {

	testRunner.Run("Should send information on finish When called", func(t *testing.T) {
		receiveDone := make(chan bool)

		useCase := getUseCaseInstance()

		testProgramName := "prname"
		go useCase.CompileProgram(testProgramName, receiveDone)

		if !<-receiveDone {
			utils.ShowError(t, true, false)
		}
	})

}
