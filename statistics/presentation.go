package statistics

import (
	"fmt"
	"go-checker/utils"
)

type Presenter struct {
	views IView
}

func ConstructPresenter(views IView) Presenter {
	return Presenter{views: views}
}

func (presenter Presenter) PresentStatisticsAppToApp(statsAppToApp *StatisticsAppToApp, totalTestsCount int) {
	presenter.presentStatistics("Success", statsAppToApp.Success, totalTestsCount)
	presenter.presentStatistics("Timeout due to brut", statsAppToApp.TimeoutBrut, totalTestsCount)
	presenter.presentStatistics("Timeout due to solution", statsAppToApp.TimeoutSolution, totalTestsCount)
	presenter.presentStatistics("Wrong answer", statsAppToApp.WrongAnswer, totalTestsCount)
}

func (presenter Presenter) presentStatistics(message string, stats Statistics, totalValue int) {
	value := stats.GetValue()
	percentageAsText := stats.CalculatePercentageAsString(totalValue)
	valueAsText := fmt.Sprintf("%v", value)

	valueAsTextWithColour := utils.Colours.Blue + valueAsText + utils.Colours.Reset
	percentageAsTextWithColour := utils.Colours.Blue + percentageAsText + utils.Colours.Reset

	finalMsg := fmt.Sprintf("\n%v: %v - %v  of all tests", message, valueAsTextWithColour, percentageAsTextWithColour)

	presenter.views.ShowStatisticsMessage(finalMsg)
}
