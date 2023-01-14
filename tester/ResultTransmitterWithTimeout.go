package tester

import "context"

type ResultTransmitterWithTimeout struct {
	sendResult         chan Result
	timeoutResultError IResultError
}

func ConstructResultTransmitterWithTimeout(sendResult chan Result, timeoutResultError IResultError) ResultTransmitterWithTimeout {
	return ResultTransmitterWithTimeout{sendResult: sendResult, timeoutResultError: timeoutResultError}
}

func (resultTransmitter ResultTransmitterWithTimeout) TransmitWithTimeout(timeoutCtx context.Context, output string) {
	select {
	case <-timeoutCtx.Done():
		resultTransmitter.sendResult <- ConstructResultWithError(resultTransmitter.timeoutResultError)
	default:
		resultTransmitter.sendResult <- ConstructResultWithSuccess(output)
	}

	close(resultTransmitter.sendResult)
}
