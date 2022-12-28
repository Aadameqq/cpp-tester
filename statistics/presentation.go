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

func (presenter Presenter) PresentStatisticsAppToApp(statsAppToApp StatisticsAppToApp, totalTestsCount int) {
	presenter.presentStatistics(statsAppToApp.Success, totalTestsCount)
	presenter.presentStatistics(statsAppToApp.TimeoutBrut, totalTestsCount)
	presenter.presentStatistics(statsAppToApp.TimeoutSolution, totalTestsCount)
	presenter.presentStatistics(statsAppToApp.WrongAnswer, totalTestsCount)
}

func (presenter Presenter) presentStatistics(stats Statistics, totalValue int) {
	value := stats.GetValue()
	percentageAsText := stats.CalculatePercentageAsString(totalValue)
	valueAsText := fmt.Sprintf("%v", value)

	valueAsTextWithColour := utils.Colours.Blue + valueAsText + utils.Colours.Reset
	percentageAsTextWithColour := utils.Colours.Blue + percentageAsText + utils.Colours.Reset

	finalMsg := fmt.Sprintf("%v - %v  of all tests\n", valueAsTextWithColour, percentageAsTextWithColour)

	presenter.views.ShowStatisticsMessage(finalMsg)
}
