package concurrency_management_core

type GenericTransmitter[T any] interface {
	Transmit(data T)
}
