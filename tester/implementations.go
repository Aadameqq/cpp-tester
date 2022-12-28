package tester

import (
	"context"
	"fmt"
	"go-checker/commands"
	"go-checker/utils"
	"io"
	"os/exec"
)

type TestInputGeneratorImpl struct {
	commands commands.Commands
}

func ConstructTestInputGeneratorImpl(commands commands.Commands) TestInputGeneratorImpl {
	return TestInputGeneratorImpl{commands: commands}
}

func (testInputGenerator TestInputGeneratorImpl) Generate(index int) string {
	command := exec.Command(testInputGenerator.commands.Run("get_tests"))

	return passDataToCommandInputAndGetOutput(fmt.Sprintf("%v", index), command)
}

type TestExecutorImpl struct {
	commands commands.Commands
}

func ConstructTestExecutorImpl(commands commands.Commands) TestExecutorImpl {
	return TestExecutorImpl{commands: commands}
}

func (testExecutor TestExecutorImpl) Execute(ctx context.Context, programName string, input string) string {
	cmd := exec.CommandContext(ctx, testExecutor.commands.Run(programName))

	return passDataToCommandInputAndGetOutput(input, cmd)
}

func passDataToCommandInputAndGetOutput(input string, command *exec.Cmd) string {
	stdin, err := command.StdinPipe()
	utils.HandleError(err)

	func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, input)
		utils.HandleError(err)
	}()

	data, err := command.Output()
	if _, isExitError := err.(*exec.ExitError); !isExitError {
		utils.HandleError(err)
	}
	return string(data)
}
