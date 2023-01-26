package adapters

import (
	"context"
	"os/exec"
	testerCore "tester-cpp/tester/core"
	textCommandsCore "tester-cpp/text-commands/core"
	"tester-cpp/utils"
	"time"
)

type ProgramTestedOutputGenerator struct {
	textCommands      textCommandsCore.TextCommands
	programName       string
	timeout           time.Duration
	resultTransmitter testerCore.TimeoutResultTransmitter
}

func (og ProgramTestedOutputGenerator) Generate(input string, sendResult chan testerCore.Result) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), og.timeout)
	defer cancel()

	textCmd := og.textCommands.Run(og.programName)
	cmd := exec.CommandContext(timeoutCtx, textCmd.GetName(), textCmd.GetParams()...)

	programExecutor := utils.ConstructProgramExecutor()
	output := programExecutor.ExecuteUsingCommandAndGetOutput(cmd, input)

	timeoutResult := testerCore.ConstructResultWithError(TestedSolutionTimeoutResultError)
	successResult := testerCore.ConstructResultWithSuccess(output)

	og.resultTransmitter.TransmitWithTimeout(sendResult, timeoutCtx, timeoutResult, successResult)
}
