package tester

import (
	"tester-cpp/commands"
	"tester-cpp/statistics"
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
	cmds := controller.commands

	testExecutor := ConstructTestExecutorImpl(cmds)
	testInputGenerator := ConstructTestInputGeneratorImpl(cmds)

	views := ConstructView()
	presenter := ConstructPresenter(views)

	useCases := ConstructTestRunner(testExecutor)

	stats := statistics.ConstructStatisticsAppToApp()
	statsViews := statistics.ConstructView()
	statsPresenters := statistics.ConstructPresenter(statsViews)
	defer statsPresenters.PresentStatisticsAppToApp(&stats, testsCount)

	modes := ConstructAppToAppMode(presenter, testInputGenerator, useCases, &stats)

	modes.RunTests(testsCount, time.Millisecond, 1000*time.Millisecond) //Todo: add config
}
