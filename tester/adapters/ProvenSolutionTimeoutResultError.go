package adapters

import (
	testerCore "tester-cpp/tester/core"
	"tester-cpp/utils"
)

var provenResultErrorMsg = *testerCore.ConstructResultMessage("Timeout due to proven solution", utils.Colours.Yellow)

var ProvenSolutionTimeoutResultError = *testerCore.ConstructResultError("proven-solution-output", provenResultErrorMsg)
