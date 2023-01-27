package concurrency_management_implementations

type GenericTransmitter[T any] struct {
	sendData chan T
}

func ConstructGenericTransmitter[T any](sendData chan T) *GenericTransmitter[T] {
	return &GenericTransmitter[T]{sendData: sendData}
}

func (gt GenericTransmitter[T]) Transmit(data T) {
	gt.sendData <- data
}
