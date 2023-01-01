package tester

import (
	"time"
)

var TestResultType = struct {
	Success int
	Timeout int
}{
	0, 1,
}

type TestResult struct {
	ResultType int
	Output     string
}

func (result TestResult) IsTimeout() bool {
	return result.ResultType == TestResultType.Timeout
}

type Test struct {
	input   string
	timeout time.Duration
}

func ConstructTest(input string, timeout time.Duration) Test {
	return Test{input: input, timeout: timeout}
}
