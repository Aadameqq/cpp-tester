package statistics

type IView interface {
	ShowStatisticsMessage(msg string)
}
type IPresenter interface {
	PresentStatistics(percentageAsText string, value int)
}

type StatisticsAppToApp struct { //TODO: does it brake ocd rule?
	Success         Statistics
	TimeoutBrut     Statistics
	TimeoutSolution Statistics
	WrongAnswer     Statistics
}

func ConstructStatisticsAppToApp() StatisticsAppToApp {
	return StatisticsAppToApp{}
}

func (presenter Presenter) PresentStatisticsAppToApp(statsAppToApp *StatisticsAppToApp, totalTestsCount int) {
	presenter.presentStatistics("Success", statsAppToApp.Success, totalTestsCount)
	presenter.presentStatistics("Timeout due to brut", statsAppToApp.TimeoutBrut, totalTestsCount)
	presenter.presentStatistics("Timeout due to solution", statsAppToApp.TimeoutSolution, totalTestsCount)
	presenter.presentStatistics("Wrong answer", statsAppToApp.WrongAnswer, totalTestsCount)
}
