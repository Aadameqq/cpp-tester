package main

import (
	"tester-cpp/commands"
	"tester-cpp/compiler"
	"tester-cpp/tester"
)

func main() {
	compilerController := compiler.ConstructAppToAppController(commands.LinuxCommands)
	compilerController.HandleCompilation()

	testerController := tester.ConstructAppToAppController(commands.LinuxCommands)
	testerController.RunTestsAppToApp()
}
