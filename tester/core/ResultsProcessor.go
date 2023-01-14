package tester_core

type ResultsProcessor interface {
	Process(firstResult Result, secondResult Result) ResultMessage
}
