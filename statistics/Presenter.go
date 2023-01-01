package statistics

import (
	"fmt"
	"go-checker/utils"
)

type Presenter struct {
	view IView
}

func ConstructPresenter(view IView) Presenter {
	return Presenter{view: view}
}

func (presenter Presenter) presentStatistics(message string, stats Statistics, totalValue int) {
	value := stats.GetValue()
	percentageAsText := stats.CalculatePercentageAsString(totalValue)
	valueAsText := fmt.Sprintf("%v", value)

	valueAsTextWithColour := utils.Colours.Blue + valueAsText + utils.Colours.Reset
	percentageAsTextWithColour := utils.Colours.Blue + percentageAsText + utils.Colours.Reset

	finalMsg := fmt.Sprintf("\n%v: %v - %v  of all tests", message, valueAsTextWithColour, percentageAsTextWithColour)

	presenter.view.ShowStatisticsMessage(finalMsg)
}
