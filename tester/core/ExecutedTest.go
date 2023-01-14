package tester_core

type ExecutedTest struct {
	input        string
	provenOutput string
	testedOutput string
}

func ConstructExecutedTest(input string, provenOutput string, testedOutput string) *ExecutedTest {
	return &ExecutedTest{input: input, provenOutput: provenOutput, testedOutput: testedOutput}
}

func (e ExecutedTest) GetInput() string {
	return e.input
}

func (e ExecutedTest) GetProvenOutput() string {
	return e.provenOutput
}

func (e ExecutedTest) GetTestedOutput() string {
	return e.testedOutput
}
