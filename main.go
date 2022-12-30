package main

import (
	"go-checker/commands"
	"go-checker/compiler"
	"go-checker/tester"
)

func main() {
	compilerController := compiler.ConstructAppToAppController(commands.LinuxCommands)
	compilerController.HandleCompilation()

	testerController := tester.ConstructAppToAppController(commands.LinuxCommands)
	testerController.RunTestsAppToApp()
}
