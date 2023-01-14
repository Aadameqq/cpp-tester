package tester_adapters

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"tester-cpp/commands"
	"tester-cpp/utils"
)

type ITestExecutor interface {
	Execute(ctx context.Context, programName string, input string) string
}

type IPresenter interface {
	PresentTimeoutBrut(index int)
	PresentTimeoutSolution(index int)
	PresentSuccess(index int)
	PresentWrongAnswer(index int, brutOutput string, solutionOutput string)
}

type IView interface {
	SetTestResultMessage(msg string)
}

type TestExecutorImpl struct {
	commands commands.Commands
}

func ConstructTestExecutorImpl(commands commands.Commands) TestExecutorImpl {
	return TestExecutorImpl{commands: commands}
}

func (testExecutor TestExecutorImpl) Execute(ctx context.Context, programName string, inputData string) string {
	cmd := exec.CommandContext(ctx, testExecutor.commands.Run(programName))

	return passDataToCommandInputAndGetOutput(inputData, cmd)
}

func passDataToCommandInputAndGetOutput(inputData string, command *exec.Cmd) string {
	inputWriter := getCommandInputWriter(command)

	writeToInputWriter(inputWriter, inputData)

	data, err := command.Output()

	if isErrorThatShouldBeHandled(err) {
		utils.HandleError(err)
	}
	return string(data)
}

func getCommandInputWriter(command *exec.Cmd) io.WriteCloser {
	inputWriter, err := command.StdinPipe()
	utils.HandleError(err)
	return inputWriter
}

func writeToInputWriter(inputWriter io.WriteCloser, inputData string) {
	defer inputWriter.Close()
	_, err := io.WriteString(inputWriter, inputData)
	utils.HandleError(err)
}

func isErrorThatShouldBeHandled(err error) bool {
	_, isExitError := err.(*exec.ExitError)
	return err != nil && !isExitError && !errors.Is(err, context.DeadlineExceeded)
}
