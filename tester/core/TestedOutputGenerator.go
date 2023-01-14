package tester_core

type TestedOutputGenerator interface {
	Generate(input string, sendOutput chan Result) Result
}
