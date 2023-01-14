package tester_core

type ProvenOutputGenerator interface {
	Generate(input string, sendOutput chan Result)
}
