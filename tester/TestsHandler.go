package tester

type ITestInputGenerator interface {
	Generate() string
}

type ITestOutputGenerator interface {
	Generate(input string, sendOutput chan string)
}

type ITestProcessor interface {
	process(input string, receiveOutput chan string) IResultPresenter
}

type IResultPresenter interface {
	Present()
}

type TestsHandler struct {
	testInputGenerator  ITestInputGenerator
	testOutputGenerator ITestOutputGenerator
	testProcessor       ITestProcessor
}

func (testsHandler TestsHandler) Handle() {
	input := testsHandler.testInputGenerator.Generate()

	receiveOutput := make(chan string)

	go testsHandler.testOutputGenerator.Generate(input, receiveOutput)

	resultPresenter := testsHandler.testProcessor.process(input, receiveOutput)

	resultPresenter.Present()

	//TODO: should add countStats here?
}
