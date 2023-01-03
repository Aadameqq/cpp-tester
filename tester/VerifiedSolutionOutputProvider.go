package tester

import "time"

type VerifiedSolutionOutputProvider struct {
	testExecutor ITestExecutor
	timeout      time.Duration
}

func (outputProvider VerifiedSolutionOutputProvider) Provide(input string, sendOutput chan string) {
	receiveOutput := make(chan Result)

	resultTransmitter := ConstructResultTransmitter(receiveOutput, VerifiedSolutionTimeoutResultError{})
	testRunner := ConstructTestRunner(outputProvider.testExecutor, resultTransmitter)

	test := ConstructTest(input, outputProvider.timeout)

	testRunner.RunWithTimeout(test, "verified_solution")

}
