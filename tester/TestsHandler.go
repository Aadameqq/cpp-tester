package tester

type IInputGenerator interface {
	Generate(index int) string
}

type IProvenOutputGenerator interface {
	Generate(input string, sendOutput chan Result)
}

type ITestedOutputGenerator interface {
	Generate(input string, sendOutput chan Result) Result
}

type IResultsProcessor interface {
	Process(firstResult Result, secondResult Result) ResultMessage
}

type IExecutedTestSaver interface {
	Save(executedTest ExecutedTest)
}

type IResultsConverter interface {
	ToExecutedTest(firstResult Result, secondResult Result, input string) (ExecutedTest, bool)
}

type IResultMessagePresenter interface {
	Present(message ResultMessage)
}

type TestsHandler struct {
	inputGenerator         IInputGenerator
	provenOutputGenerator  IProvenOutputGenerator
	testedOutputGenerator  ITestedOutputGenerator
	executedTestSaver      IExecutedTestSaver
	resultsProcessor       IResultsProcessor
	resultsConverter       IResultsConverter
	resultMessagePresenter IResultMessagePresenter
}

func (th TestsHandler) Handle(testsCount int) {
	for i := 0; i < testsCount; i++ {
		th.handleSingle(i)
	}
}

func (th TestsHandler) handleSingle(index int) {
	input := th.inputGenerator.Generate(index)
	provenResult, testedResult := th.generateOutputs(input)
	finalResultMessage := th.processResults(provenResult, testedResult)
	th.resultMessagePresenter.Present(finalResultMessage)
	th.save(provenResult, testedResult, input)
}

func (th TestsHandler) processResults(provenResult Result, testedResult Result) ResultMessage {
	return th.resultsProcessor.Process(provenResult, testedResult)
}

func (th TestsHandler) save(provenResult Result, testedResult Result, input string) {
	executedTest, isError := th.resultsConverter.ToExecutedTest(provenResult, testedResult, input)
	if !isError {
		go th.executedTestSaver.Save(executedTest)
	}
}

func (th TestsHandler) generateOutputs(input string) (Result, Result) {
	receiveProvenResult := make(chan Result)
	receiveTestedResult := make(chan Result)

	go th.provenOutputGenerator.Generate(input, receiveProvenResult)
	go th.testedOutputGenerator.Generate(input, receiveTestedResult)

	provenResult := <-receiveProvenResult
	testedResult := <-receiveTestedResult

	return provenResult, testedResult
}
