package tester

type VerifiedSolutionTimeoutResultError struct {
}

func (err VerifiedSolutionTimeoutResultError) GetName() string {
	return "timeout"
}

func (err VerifiedSolutionTimeoutResultError) GetMessage() string {
	return "Timeout due to verified solution"
}

func (err VerifiedSolutionTimeoutResultError) GetColour() string {
	return ""
}
