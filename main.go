package main

import (
	"go-checker/commands"
	"go-checker/tester"
)

func main() {
	controller := tester.ConstructAppToAppController(commands.LinuxCommands)
	controller.RunTestsAppToApp()
}
