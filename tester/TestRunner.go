package tester

import "context"

type TestRunner struct {
	testExecutor ITestExecutor
}

func ConstructTestRunner(testExecutor ITestExecutor) TestRunner {
	return TestRunner{testExecutor: testExecutor}
}

func (testRunner TestRunner) Run(test Test, result chan TestResult, programName string) {
	ctx, cancel := context.WithTimeout(context.Background(), test.timeout)
	defer cancel()

	output := testRunner.testExecutor.Execute(ctx, programName, test.input)

	select {
	case <-ctx.Done(): //TODO: should it be somewhere else?
		result <- TestResult{ResultType: TestResultType.Timeout}
	default:
		result <- TestResult{ResultType: TestResultType.Success, Output: output}
	}

	close(result)
}
