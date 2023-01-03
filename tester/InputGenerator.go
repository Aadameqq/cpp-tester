package tester

import (
	"fmt"
	"os/exec"
	"tester-cpp/commands"
)

type InputGenerator struct {
	commands commands.Commands
}

func ConstructTestInputGeneratorImpl(commands commands.Commands) InputGenerator {
	return InputGenerator{commands: commands}
}

func (inputGenerator InputGenerator) Generate(index int) string {
	command := exec.Command(inputGenerator.commands.Run("get_tests"))

	return passDataToCommandInputAndGetOutput(fmt.Sprintf("%v", index), command)
	// Todo: "passDataToCommandInputAndGetOutput" move this function declaration somewhere else
}
