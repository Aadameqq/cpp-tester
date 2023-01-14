package tester_core

type TResultsConverter struct {
}

func ConstructResultsConverter() *TResultsConverter {
	return &TResultsConverter{}
}

func (T TResultsConverter) ToExecutedTest(provenResult Result, testedResult Result, input string) (executedTest ExecutedTest, isError bool) {
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
