package utils

import (
	"context"
	"errors"
	"io"
	"os/exec"
)

type ProgramExecutor struct {
}

func ConstructProgramExecutor() *ProgramExecutor {
	return &ProgramExecutor{}
}

func (pe ProgramExecutor) ExecuteUsingCommandAndGetOutput(programExecCommand *exec.Cmd, inputData string) string {
	inputWriter := pe.getCommandInputWriter(programExecCommand)

	pe.writeToInput(inputWriter, inputData)

	data, err := programExecCommand.Output()

	if pe.isErrorThatShouldBeHandled(err) {
		HandleError(err)
	}

	return string(data)
}

func (pe ProgramExecutor) getCommandInputWriter(command *exec.Cmd) io.WriteCloser {
	inputWriter, err := command.StdinPipe()
	HandleError(err)
	return inputWriter
}

func (pe ProgramExecutor) writeToInput(inputWriter io.WriteCloser, inputData string) {
	defer inputWriter.Close()
	_, err := io.WriteString(inputWriter, inputData)
	HandleError(err)
}

func (pe ProgramExecutor) isErrorThatShouldBeHandled(err error) bool {
	_, isExitError := err.(*exec.ExitError)
	return err != nil && !isExitError && !errors.Is(err, context.DeadlineExceeded)
}
