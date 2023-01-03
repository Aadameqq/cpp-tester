package tester

type IResultError interface {
	GetName() string
	GetMessage() string
	GetColour() string
}

type NoneResultError struct {
}

func (err NoneResultError) GetColour() string {
	return ""
}

func (err NoneResultError) GetName() string {
	return "none"
}

func (err NoneResultError) GetMessage() string {
	return "None"
}

type Result struct {
	resultError  IResultError
	payload      string
	isSuccessful bool
}

func ConstructResultWithError(resultError IResultError) Result {
	return Result{resultError: resultError}
}

func ConstructResultWithSuccess(payload string) Result {
	return Result{isSuccessful: true, payload: payload}
}

func (result *Result) IsSuccessful() bool {
	return result.isSuccessful
}

func (result *Result) GetPayloadIfSuccessful() *string {
	if result.IsSuccessful() {
		return &result.payload
	}
	return nil
}

func (result *Result) IsError() bool {
	return result.resultError.GetName() != NoneResultError{}.GetName()
}

func (result *Result) IsErrorType(resultType IResultError) bool {
	return result.resultError.GetName() == resultType.GetName()
}

func (result *Result) GetErrorIfIsError() *IResultError {
	if result.IsError() {
		return &result.resultError
	}
	return nil
}
