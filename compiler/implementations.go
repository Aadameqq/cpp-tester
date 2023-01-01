package compiler

import (
	"os/exec"
	"strings"
	"tester-cpp/commands"
	"tester-cpp/utils"
)

type CompilationRunnerImpl struct {
	commands commands.Commands
}

func ConstructCompilationRunnerImpl(commands commands.Commands) CompilationRunnerImpl {
	return CompilationRunnerImpl{commands: commands}
}

func (compilationRunner CompilationRunnerImpl) Run(programName string) {
	commandName, commandParams := compilationRunner.getCommandNameAndItsParameters(programName)
	command := exec.Command(commandName, commandParams...)
	compilationRunner.runCommand(command)
}

func (compilationRunner CompilationRunnerImpl) runCommand(command *exec.Cmd) {
	utils.HandleError(command.Run())
}

func (compilationRunner CompilationRunnerImpl) getCommandNameAndItsParameters(programName string) (string, []string) {
	splitCommand := strings.Split(compilationRunner.commands.Compile(programName), " ")
	return splitCommand[0], splitCommand[1:]
}
