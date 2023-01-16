package commands_implementations

import "fmt"

type CppLinuxCommands struct {
}

func (lc CppLinuxCommands) Compile(programName string) string {
	return "/usr/bin/g++ -o " + programName + ".out " + programName + ".cpp"
}

func (lc CppLinuxCommands) Run(programName string) string {
	return fmt.Sprintf("./%v.out", programName)
}
