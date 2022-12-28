package tester

import "fmt"

type View struct {
}

func ConstructView() View {
	return View{}
}

func (view View) setTestStatusMessage(msg string) {
	fmt.Println(msg)
}
