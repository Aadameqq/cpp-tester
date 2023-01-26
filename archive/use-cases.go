package archive

type UseCase struct {
	compilationRunner ICompilationRunner
}

func ConstructUseCase(compilationRunner ICompilationRunner) UseCase {
	return UseCase{compilationRunner: compilationRunner}
}

func (useCase UseCase) CompileProgram(programName string, sendDone chan bool) {
	useCase.compilationRunner.Run(programName)
	sendDone <- true
}
