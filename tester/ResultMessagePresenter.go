package tester

import (
	"fmt"
)

type ResultMessagePresenter struct {
	view IView
}

func ConstructResultMessagePresenter(view IView) ResultMessagePresenter {
	return ResultMessagePresenter{view: view}
}

func (presenter ResultMessagePresenter) PresentMessage(index int, msg string) string {
	return fmt.Sprintf("UnexecutedTest nr. %v - %v", index, msg)
}
