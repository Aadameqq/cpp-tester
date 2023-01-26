package archive

type AppToAppMode struct {
	useCase   IUseCase
	presenter IPresenter
}

func ConstructAppToAppMode(useCase IUseCase, presenter IPresenter) AppToAppMode {
	return AppToAppMode{useCase: useCase, presenter: presenter}
}

func (mode AppToAppMode) Compile() {
	compileProgram := mode.useCase.CompileProgram

	receiveDone := make(chan bool)
	sendDone := make(chan bool)

	go mode.presenter.PresentCompilation(sendDone)

	getTestsProgramName := "get_tests" //TODO: use config
	go compileProgram(getTestsProgramName, receiveDone)

	solutionProgramName := "solution" //TODO: use config
	go compileProgram(solutionProgramName, receiveDone)

	brutProgramName := "brut" //TODO: use config
	go compileProgram(brutProgramName, receiveDone)

	<-receiveDone
	<-receiveDone
	<-receiveDone

	sendDone <- true

	mode.presenter.PresentCompilationSuccess()
}
