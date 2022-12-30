package compiler

type ICompilationRunner interface {
	Run(programName string)
}

type IPresenter interface {
	PresentCompilation(receiveDone chan bool)
	PresentCompilationSuccess()
}

type IView interface {
	updateAnimation(msg string)
	setAnimationEndMessage(msg string)
}

type IUseCase interface {
	CompileProgram(programName string, sendDone chan bool)
}
