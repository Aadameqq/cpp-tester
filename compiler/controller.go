package compiler

import "go-checker/commands"

type AppToAppController struct {
	commands commands.Commands
}

func ConstructAppToAppController(commands commands.Commands) AppToAppController {
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
