package tester_core

type ResultsConverter interface {
	ToExecutedTest(firstResult Result, secondResult Result, input string) (ExecutedTest, bool)
}
