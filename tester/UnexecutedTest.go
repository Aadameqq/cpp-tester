package tester

import "time"

type UnexecutedTest struct {
	input   string
	timeout time.Duration
}

func ConstructUnexecutedTest(input string, timeout time.Duration) UnexecutedTest {
	return UnexecutedTest{input: input, timeout: timeout}
}

func (ut UnexecutedTest) GetInput() string {
	return ut.input
}

func (ut UnexecutedTest) GetTimeout() time.Duration {
	return ut.timeout
}
