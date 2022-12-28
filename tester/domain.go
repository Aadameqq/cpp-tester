package tester

import (
	"context"
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

type ITestExecutor interface {
	Execute(ctx context.Context, programName string, input string) string
}

type IPresenter interface {
	PresentTimeoutBrut(index int)
	PresentTimeoutSolution(index int)
	PresentSuccess(index int)
	PresentWrongAnswer(index int, brutOutput string, solutionOutput string)
}

type ITestInputGenerator interface {
	Generate(index int) string
}

type IViews interface {
	setTestStatusMessage(msg string)
}
