package compiler_core

type CompilationExecutor[ProgramLocation any] interface {
	execute(programLocation ProgramLocation)
}
