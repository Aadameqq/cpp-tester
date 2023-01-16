package tester_adapters

import (
	"fmt"
	"os/exec"
)

type InputGenerator struct {
	commands core.Commands
}

func ConstructTestInputGeneratorImpl(commands core.Commands) InputGenerator {
	return InputGenerator{commands: commands}
}

func (inputGenerator InputGenerator) Generate(index int) string {
	command := exec.Command(inputGenerator.commands.Run("get_tests"))

	return adapters.passDataToCommandInputAndGetOutput(fmt.Sprintf("%v", index), command)
	// Todo: "passDataToCommandInputAndGetOutput" move this function declaration somewhere else
}
