package adapters

import (
	"os/exec"
	textCommandsCore "tester-cpp/text-commands/core"
	"tester-cpp/utils"
)

type ProgramInputGenerator struct {
	programName  string
	textCommands textCommandsCore.TextCommands
}

func ConstructProgramInputGenerator(programName string, textCommands textCommandsCore.TextCommands) *ProgramInputGenerator {
	return &ProgramInputGenerator{programName: programName, textCommands: textCommands}
}

func (ig ProgramInputGenerator) Generate(index int) string {
	textCmd := ig.textCommands.Run(ig.programName)
	command := ig.getCommand(textCmd)
	programExecutor := utils.ConstructProgramExecutor()
	return programExecutor.ExecuteUsingCommandAndGetOutput(command, utils.ConvertIntToString(index))
}

func (ig ProgramInputGenerator) getCommand(runCmd textCommandsCore.TextCommand) *exec.Cmd {
	return exec.Command(runCmd.GetName(), runCmd.GetParams()...)
}
