package adapters

import (
	testerCore "tester-cpp/tester/core"
	"tester-cpp/utils"
)

var testedResultErrorMsg = *testerCore.ConstructResultMessage("Timeout due to proven solution", utils.Colours.Yellow)

var TestedSolutionTimeoutResultError = *testerCore.ConstructResultError("proven-solution-output", testedResultErrorMsg)
