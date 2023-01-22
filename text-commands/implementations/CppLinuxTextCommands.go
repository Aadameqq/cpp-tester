package text_commands_implementations

import (
	"fmt"
	textCommandsCore "tester-cpp/text-commands/core"
)

type CppLinuxCommands struct {
}

func (lc CppLinuxCommands) Compile(programName string) textCommandsCore.TextCommand {
	params := []string{"-o", programName + ".out", programName + ".cpp"}
	return *textCommandsCore.ConstructTextCommand("/usr/bin/g++", params)
}

func (lc CppLinuxCommands) Run(programName string) textCommandsCore.TextCommand {
	cmdName := fmt.Sprintf("./%v.out", programName)
	return *textCommandsCore.ConstructTextCommand(cmdName, []string{})
}
