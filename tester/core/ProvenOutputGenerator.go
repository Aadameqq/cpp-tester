package tester_core

type ProvenOutputGenerator interface {
	Generate(input string, sendResult chan Result)
}
