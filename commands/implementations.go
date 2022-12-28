package commands

import "fmt"

var LinuxCommands = Commands{Compile: func(programName string) string {
	return "/usr/bin/g++ -o  " + programName + ".out " + programName + ".cpp"
}, Run: func(programName string) string {
	return fmt.Sprintf("./%v.out", programName)
}}
