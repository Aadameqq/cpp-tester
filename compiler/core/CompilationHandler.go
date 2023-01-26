package compiler_core

type CompilationHandler[ProgramLocation any] struct {
	compilationPresenter CompilationPresenter
	compiler             Compiler[ProgramLocation]
}

func (ch CompilationHandler[ProgramLocation]) Handle(programLocations []ProgramLocation) {
	receiveDone := make(chan bool)
	sendDone := make(chan bool)

	go ch.compilationPresenter.Present(sendDone)

	for _, programLocation := range programLocations {
		go ch.compiler.Compile(programLocation, receiveDone)
	}

	for _, _ = range programLocations {
		<-receiveDone
	}

	sendDone <- true

	ch.compilationPresenter.PresentFinish()
}
