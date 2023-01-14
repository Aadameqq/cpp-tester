package tester_core

import (
	"context"
	testerAdapters "tester-cpp/tester/adapters"
)

type IResultTransmitter interface {
	TransmitWithTimeout(timeoutCtx context.Context, output string)
}

type TestRunner struct {
	testExecutor      testerAdapters.ITestExecutor
	resultTransmitter IResultTransmitter
}

func ConstructTestRunner(testExecutor testerAdapters.ITestExecutor, resultTransmitter IResultTransmitter) TestRunner {
	return TestRunner{testExecutor: testExecutor, resultTransmitter: resultTransmitter}
}

func (testRunner TestRunner) RunWithTimeout(test UnexecutedTest, programName string) {
	ctx, cancel := context.WithTimeout(context.Background(), test.GetTimeout())
	defer cancel()

	output := testRunner.testExecutor.Execute(ctx, programName, test.GetInput()) //TODO: will it be working properly

	testRunner.resultTransmitter.TransmitWithTimeout(ctx, output)
}
