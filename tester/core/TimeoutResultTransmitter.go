package tester_core

import "context"

type TimeoutResultTransmitter interface {
	TransmitWithTimeout(sendResult chan Result, timeoutCtx context.Context, timeoutResult Result, noTimeoutResult Result)
}
