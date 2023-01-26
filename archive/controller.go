package archive

type AppToAppController struct {
	commands core.Commands
}

func ConstructAppToAppController(commands core.Commands) AppToAppController {
	return AppToAppController{commands: commands}
}

func (controller AppToAppController) HandleCompilation() {
	view := ConstructView()
	presenter := ConstructPresenter(view)
	compilationRunner := ConstructCompilationRunnerImpl(controller.commands)
	useCase := ConstructUseCase(compilationRunner)
	mode := ConstructAppToAppMode(useCase, presenter)
	mode.Compile()
}
