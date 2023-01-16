package commands_core

type Commands interface {
	Compile(programName string) string
	Run(programName string) string
}
