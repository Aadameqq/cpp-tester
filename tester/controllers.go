package tester

import (
	"go-checker/commands"
	"go-checker/statistics"
	"time"
)

type AppToAppController struct {
	commands commands.Commands
	//Todo: add config
}

func ConstructAppToAppController(commands commands.Commands) AppToAppController {
	return AppToAppController{commands: commands}
}

func (controller AppToAppController) RunTestsAppToApp() {
	testsCount := 240
	commands := controller.commands

	testExecutor := ConstructTestExecutorImpl(commands)
	testInputGenerator := ConstructTestInputGeneratorImpl(commands)

	views := ConstructView()
	presenter := ConstructPresenter(views)

	useCases := ConstructUseCase(testExecutor)

	stats := statistics.ConstructStatisticsAppToApp()
	statsViews := statistics.ConstructView()
	statsPresenters := statistics.ConstructPresenter(statsViews)
	defer statsPresenters.PresentStatisticsAppToApp(stats, testsCount)

	modes := ConstructAppToAppMode(presenter, testInputGenerator, useCases, stats)

	modes.RunTests(testsCount, 50*time.Millisecond, 1000*time.Millisecond) //Todo: add config
}
