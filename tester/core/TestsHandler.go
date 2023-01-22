package tester_core

type TestsHandler struct {
	inputGenerator         InputGenerator
	provenOutputGenerator  ProvenOutputGenerator
	testedOutputGenerator  TestedOutputGenerator
	executedTestRepository ExecutedTestRepository
	resultsProcessor       ResultsProcessor
	resultMessagePresenter ResultMessagePresenter
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
	resultsConverter := ConstructResultsConverter()
	executedTest, isError := resultsConverter.ToExecutedTest(provenResult, testedResult, input)
	if !isError {
		go th.executedTestRepository.Save(executedTest)
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
