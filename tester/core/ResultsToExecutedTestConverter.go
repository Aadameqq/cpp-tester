package tester_core

type TResultsToExecutedTestConverter struct {
}

func ConstructResultsToExecutedTestConverter() *TResultsToExecutedTestConverter {
	return &TResultsToExecutedTestConverter{}
}

func (T TResultsToExecutedTestConverter) Convert(provenResult Result, testedResult Result, input string) (executedTest ExecutedTest, isError bool) {
	provenOutput, conditionsMet := provenResult.GetPayloadIfSuccess()

	if !conditionsMet {
		return ExecutedTest{}, true
	}

	testedOutput, conditionsMet := testedResult.GetPayloadIfSuccess()

	if !conditionsMet {
		return ExecutedTest{}, true
	}

	return *ConstructExecutedTest(input, provenOutput, testedOutput), false
}
