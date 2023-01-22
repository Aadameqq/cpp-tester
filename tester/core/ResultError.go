package tester_core

type ResultError struct {
	name    string
	message ResultMessage
}

func ConstructResultError(name string, message ResultMessage) *ResultError {
	return &ResultError{name: name, message: message}
}

func (re ResultError) GetName() string {
	return re.name
}

func (re ResultError) GetMessage() ResultMessage {
	return re.message
}
