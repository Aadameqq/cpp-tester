package tester_core

type ResultError interface {
	GetName() string
	GetMessage() ResultMessage
} //TODO: shouldn't it be an object?
