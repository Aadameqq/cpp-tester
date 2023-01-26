package compiler_core

type CompilationPresenter interface {
	Present(receiveDone chan bool)
	PresentFinish()
}
