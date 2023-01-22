package tester_core

type ExecutedTestRepository interface {
	Save(executedTest ExecutedTest)
}
