package tester_core

type TestedOutputGenerator interface {
	Generate(input string, sendResult chan Result)
}
