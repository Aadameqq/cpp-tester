package tester

import "context"

type TestRunner struct {
	testExecutor ITestExecutor
}

func ConstructTestRunner(testExecutor ITestExecutor) TestRunner {
	return TestRunner{testExecutor: testExecutor}
}

func (testRunner TestRunner) RunWithTimeout(test Test, result chan TestResult, programName string) {
	ctx, cancel := context.WithTimeout(context.Background(), test.timeout)
	defer cancel()

	output := testRunner.testExecutor.Execute(ctx, programName, test.input)

	testRunner.sendResult(ctx, result, output)
}

func (testRunner TestRunner) sendResult(ctx context.Context, result chan TestResult, output string) {
	select {
	case <-ctx.Done():
		result <- TestResult{ResultType: TestResultType.Timeout}
	default:
		result <- TestResult{ResultType: TestResultType.Success, Output: output}
	}

	close(result)
}
