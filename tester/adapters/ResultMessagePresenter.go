package tester_adapters

import (
	"fmt"
)

type ResultMessagePresenter struct {
	view adapters.IView
}

func ConstructResultMessagePresenter(view adapters.IView) ResultMessagePresenter {
	return ResultMessagePresenter{view: view}
}

func (presenter ResultMessagePresenter) PresentMessage(index int, msg string) string {
	return fmt.Sprintf("UnexecutedTest nr. %v - %v", index, msg)
}
