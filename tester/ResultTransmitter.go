package tester

import "context"

type ResultTransmitter struct {
	sendResult         chan Result
	timeoutResultError IResultError
}

func ConstructResultTransmitter(sendResult chan Result, timeoutResultError IResultError) ResultTransmitter {
	return ResultTransmitter{sendResult: sendResult, timeoutResultError: timeoutResultError}
}

func (resultTransmitter ResultTransmitter) TransmitWithTimeout(timeoutCtx context.Context, output string) {
	select {
	case <-timeoutCtx.Done():
		resultTransmitter.sendResult <- ConstructResultWithError(resultTransmitter.timeoutResultError)
	default:
		resultTransmitter.sendResult <- ConstructResultWithSuccess(output)
	}

	close(resultTransmitter.sendResult)
}
