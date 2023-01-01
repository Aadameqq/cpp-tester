package tester

import "context"

type ITestExecutor interface {
	Execute(ctx context.Context, programName string, input string) string
}

type IPresenter interface {
	PresentTimeoutBrut(index int)
	PresentTimeoutSolution(index int)
	PresentSuccess(index int)
	PresentWrongAnswer(index int, brutOutput string, solutionOutput string)
}

type ITestInputGeneratorTemporary interface {
	Generate(index int) string
}

type IView interface {
	setTestStatusMessage(msg string)
}
