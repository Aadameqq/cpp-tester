package tester_core

type Result struct {
	error     ResultError
	payload   string
	isSuccess bool
}

func ConstructResultWithError(resultError ResultError) Result {
	return Result{error: resultError}
}

func ConstructResultWithSuccess(payload string) Result {
	return Result{isSuccess: true, payload: payload}
}

func (result *Result) IsSuccess() bool {
	return result.isSuccess
}

func (result *Result) GetPayloadIfSuccess() (payload string, conditionsMet bool) {
	if result.IsSuccess() {
		return result.payload, true
	}
	return "", false
}

func (result *Result) IsError() bool {
	return result.error == ResultError{} //TODO: check if that works
}

func (result *Result) IsErrorInstanceOf(resultError ResultError) bool {
	if !result.IsError() {
		return false
	}
	return result.error.GetName() == resultError.GetName()
}

func (result *Result) GetErrorIfError() (resultError ResultError, conditionsMet bool) {
	if result.IsError() {
		return result.error, true
	}
	return ResultError{}, false
}
