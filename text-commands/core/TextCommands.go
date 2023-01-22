package text_commands_core

type TextCommands interface {
	Compile(programName string) TextCommand
	Run(programName string) TextCommand
}
