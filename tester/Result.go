package tester

type IResultError interface {
	GetName() string
	GetMessage() ResultMessage
}

type Result struct {
	error     IResultError
	payload   string
	isSuccess bool
}

func ConstructResultWithError(resultError IResultError) Result {
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
	return result.error == nil
}

func (result *Result) IsErrorInstanceOf(resultError IResultError) bool {
	if !result.IsError() {
		return false
	}
	return result.error.GetName() == resultError.GetName()
}

func (result *Result) GetErrorIfError() (resultError IResultError, conditionsMet bool) {
	if result.IsError() {
		return result.error, true
	}
	return nil, false
}
