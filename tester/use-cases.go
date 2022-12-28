package tester

import "context"

type UseCase struct {
	testExecutor ITestExecutor
}

func ConstructUseCase(testExecutor ITestExecutor) UseCase {
	return UseCase{testExecutor: testExecutor}
}

func (useCase UseCase) RunTest(test Test, result chan TestResult, programName string) {
	ctx, cancel := context.WithTimeout(context.Background(), test.timeout)
	defer cancel()

	output := useCase.testExecutor.Execute(ctx, programName, test.input)

	select {
	case <-ctx.Done():
		result <- TestResult{ResultType: TestResultType.Timeout}
	default:
		result <- TestResult{ResultType: TestResultType.Success, Output: output}
	}

	close(result)
}
