package tester

type IInputProvider interface {
	Provide(index int) string
}

type IOutputProvider interface {
	Provide(input string, sendOutput chan Result)
}

type ITestProcessor interface {
	process(input string, receiveOutput chan Result) Result
}

type IResultPresenter interface {
	Present(result Result)
}

type TestsHandler struct {
	inputProvider   IInputProvider
	outputProvider  IOutputProvider
	testProcessor   ITestProcessor
	resultPresenter IResultPresenter
}

func (testsHandler TestsHandler) Handle() {
	index := 1 //TODO: add loop
	input := testsHandler.inputProvider.Provide(index)

	receiveOutput := make(chan Result)

	go testsHandler.outputProvider.Provide(input, receiveOutput)

	result := testsHandler.testProcessor.process(input, receiveOutput) //TODO: should it be divided into two classes

	testsHandler.resultPresenter.Present(result)

	//TODO: should add countStats here?
}
