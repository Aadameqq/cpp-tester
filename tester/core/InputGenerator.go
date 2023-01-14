package tester_core

type InputGenerator interface {
	Generate(index int) string
}
