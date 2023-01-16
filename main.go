package main

import (
	"tester-cpp/compiler"
	"tester-cpp/tester"
)

func main() {
	compilerController := compiler.ConstructAppToAppController(implementations.LinuxCommands)
	compilerController.HandleCompilation()

	testerController := tester.ConstructAppToAppController(implementations.LinuxCommands)
	testerController.RunTestsAppToApp()
}
