package tester

import "context"

type ResultTransmitter struct {
	result chan TestResult
}

func (resultTransmitter ResultTransmitter) TransmitWithTimeout(timeoutCtx context.Context, output string) {
	select {
	case <-timeoutCtx.Done():
		resultTransmitter.result <- TestResult{ResultType: TestResultType.Timeout}
	default:
		resultTransmitter.result <- TestResult{ResultType: TestResultType.Success, Output: output}
	}

	close(resultTransmitter.result)
}
