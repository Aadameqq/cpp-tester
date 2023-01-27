package compiler_core

import concurrencyCore "tester-cpp/concurrency-management/core"

type Compiler[ProgramLocation any] struct {
	compilationExecutor CompilationExecutor[ProgramLocation]
	doneTransmitter     concurrencyCore.GenericTransmitter[bool]
}

func (c Compiler[ProgramLocation]) Compile(programLocation ProgramLocation) {
	c.compilationExecutor.execute(programLocation)
	c.doneTransmitter.Transmit(true)
}
