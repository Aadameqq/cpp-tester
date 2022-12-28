package commands

type Commands struct {
	Compile func(programName string) string
	Run     func(programName string) string
}
