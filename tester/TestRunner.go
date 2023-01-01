package tester

import "context"

type IResultTransmitter interface {
	TransmitWithTimeout(timeoutCtx context.Context, output string)
}

type TestRunner struct {
	testExecutor      ITestExecutor
	resultTransmitter IResultTransmitter
}

func ConstructTestRunner(testExecutor ITestExecutor, resultTransmitter IResultTransmitter) TestRunner {
	return TestRunner{testExecutor: testExecutor, resultTransmitter: resultTransmitter}
}

func (testRunner TestRunner) RunWithTimeout(test Test, programName string) {
	ctx, cancel := context.WithTimeout(context.Background(), test.timeout)
	defer cancel()

	output := testRunner.testExecutor.Execute(ctx, programName, test.input) //TODO: maybe remove programName

	testRunner.resultTransmitter.TransmitWithTimeout(ctx, output)
}
