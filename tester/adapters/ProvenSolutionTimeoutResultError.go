package adapters

import (
	testerCore "tester-cpp/tester/core"
	"tester-cpp/utils"
)

var message = *testerCore.ConstructResultMessage("Timeout due to proven solution", utils.Colours.Yellow)

var ProvenSolutionTimeoutResultError = *testerCore.ConstructResultError("proven-solution-output", message)
