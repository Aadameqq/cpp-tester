package statistics

import "fmt"

type View struct {
}

func ConstructView() View {
	return View{}
}

func (views View) ShowStatisticsMessage(msg string) {
	fmt.Println(msg)
}
