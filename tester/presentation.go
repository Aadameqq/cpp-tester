package tester

import (
	"fmt"
	"go-checker/utils"
)

type Presenter struct {
	views IViews
}

func ConstructPresenter(views IViews) Presenter {
	return Presenter{views: views}
}

func (presenter Presenter) PresentTimeoutBrut(index int) {
	msg := presenter.presentResponse(index, utils.Colours.Yellow+"Timeout due to brut"+utils.Colours.Reset)
	presenter.views.setTestStatusMessage(msg)
}

func (presenter Presenter) PresentTimeoutSolution(index int) {
	msg := presenter.presentResponse(index, utils.Colours.Yellow+"Timeout due to solution"+utils.Colours.Reset)
	presenter.views.setTestStatusMessage(msg)
}

func (presenter Presenter) PresentSuccess(index int) {
	msg := presenter.presentResponse(index, utils.Colours.Green+"Success"+utils.Colours.Reset)
	presenter.views.setTestStatusMessage(msg)
}

func (presenter Presenter) PresentWrongAnswer(index int, brutOutput string, solutionOutput string) {
	msg := presenter.presentResponse(index, utils.Colours.Red+"Wrong answer"+utils.Colours.Reset+"\n Brut:"+brutOutput+"\n\nSolution:"+solutionOutput)
	presenter.views.setTestStatusMessage(msg)
}

func (presenter Presenter) presentResponse(index int, msg string) string {
	return fmt.Sprintf("Test nr. %v - %v", index, msg)
}
