package tester

import (
	"go-checker/statistics"
	"time"
)

type AppToAppMode struct {
	presenter          IPresenter
	testInputGenerator ITestInputGenerator
	useCase            UseCase
	stats              *statistics.StatisticsAppToApp
}

func ConstructAppToAppMode(presenter IPresenter, testInputGenerator ITestInputGenerator, useCases UseCase, stats *statistics.StatisticsAppToApp) AppToAppMode {
	return AppToAppMode{presenter: presenter, testInputGenerator: testInputGenerator, useCase: useCases, stats: stats}
}

func (mode AppToAppMode) RunTests(testsCount int, brutTimeout time.Duration, solutionTimeout time.Duration) {
	for i := 0; i < testsCount; i++ {
		mode.runTest(i, brutTimeout, solutionTimeout)
	}
}

func (mode AppToAppMode) runTest(index int, brutTimeout time.Duration, solutionTimeout time.Duration) {
	testInput := mode.testInputGenerator.Generate(index)

	brutChannel := make(chan TestResult)
	solutionChannel := make(chan TestResult)

	brutTest := ConstructTest(testInput, brutTimeout)
	solutionTest := ConstructTest(testInput, solutionTimeout)

	go mode.useCase.RunTest(brutTest, brutChannel, "brut") //TODO: remove program names
	go mode.useCase.RunTest(solutionTest, solutionChannel, "solution")

	brutTestResult := <-brutChannel
	solutionTestResult := <-solutionChannel

	if brutTestResult.IsTimeout() {
		mode.handleBrutTimeout(index)
		return
	}

	if solutionTestResult.IsTimeout() {
		mode.handleSolutionTimeout(index)
		return
	}

	if brutTestResult.Output == solutionTestResult.Output {
		mode.handleSuccess(index)
	} else {
		mode.handleWrongAnswer(index, brutTestResult, solutionTestResult)
	}
}

func (mode AppToAppMode) handleWrongAnswer(index int, brutTestResult TestResult, solutionTestResult TestResult) {
	mode.presenter.PresentWrongAnswer(index, brutTestResult.Output, solutionTestResult.Output)
	mode.stats.WrongAnswer.IncreaseValue()
}

func (mode AppToAppMode) handleSuccess(index int) {
	mode.presenter.PresentSuccess(index)
	mode.stats.Success.IncreaseValue()
}

func (mode AppToAppMode) handleSolutionTimeout(index int) {
	mode.presenter.PresentTimeoutSolution(index)
	mode.stats.TimeoutSolution.IncreaseValue()
}

func (mode AppToAppMode) handleBrutTimeout(index int) {
	mode.presenter.PresentTimeoutBrut(index)
	mode.stats.TimeoutBrut.IncreaseValue()
}
