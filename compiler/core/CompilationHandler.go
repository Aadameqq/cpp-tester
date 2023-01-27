package compiler_core

type CompilationHandler[ProgramLocation any] struct {
	compilationPresenter CompilationPresenter
	compiler             Compiler[ProgramLocation]
	receiveDone          chan bool
}

func (ch CompilationHandler[ProgramLocation]) Handle(programLocations []ProgramLocation) {
	sendDone := make(chan bool)

	go ch.compilationPresenter.Present(sendDone)

	for _, programLocation := range programLocations {
		go ch.compiler.Compile(programLocation)
	}

	for _, _ = range programLocations {
		<-ch.receiveDone
	}

	sendDone <- true

	ch.compilationPresenter.PresentFinish()
}
