package compiler_core

type Compiler[ProgramLocation any] struct {
	compilationExecutor CompilationExecutor[ProgramLocation]
}

func (c Compiler[ProgramLocation]) Compile(programLocation ProgramLocation, sendDone chan bool) {
	c.compilationExecutor.execute(programLocation)
	sendDone <- true //TODO: separate thread logic
}
