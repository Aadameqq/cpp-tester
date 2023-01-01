package tester

type IInputProvider interface {
	Provide(index int) string
}

type IOutputProvider interface {
	Provide(input string, sendOutput chan string)
}

type ITestProcessor interface {
	process(input string, receiveOutput chan string) IResultPresenter
}

type IResultPresenter interface {
	Present()
}

type TestsHandler struct {
	inputProvider  IInputProvider
	outputProvider IOutputProvider
	testProcessor  ITestProcessor
}

func (testsHandler TestsHandler) Handle() {
	index := 1 //TODO: add loop
	input := testsHandler.inputProvider.Provide(index)

	receiveOutput := make(chan string)

	go testsHandler.outputProvider.Provide(input, receiveOutput)

	resultPresenter := testsHandler.testProcessor.process(input, receiveOutput)

	resultPresenter.Present()

	//TODO: should add countStats here?
}
