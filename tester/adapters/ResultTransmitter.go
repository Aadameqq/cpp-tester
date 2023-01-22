package adapters

import (
	"context"
	testerCore "tester-cpp/tester/core"
)

type ResultTransmitter struct {
}

func ConstructResultTransmitter() *ResultTransmitter {
	return &ResultTransmitter{}
}

func (rt ResultTransmitter) TransmitWithTimeout(sendResult chan testerCore.Result, timeoutCtx context.Context, timeoutResult testerCore.Result, noTimeoutResult testerCore.Result) {
	select {
	case <-timeoutCtx.Done():
		sendResult <- timeoutResult
	default:
		sendResult <- noTimeoutResult
	}

	close(sendResult)
}
