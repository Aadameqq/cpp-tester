package tester

import "time"

type Test struct {
	input   string
	timeout time.Duration
}

func ConstructTest(input string, timeout time.Duration) Test {
	return Test{input: input, timeout: timeout}
}
