package tester_adapters

import "fmt"

type View struct {
}

func ConstructView() View {
	return View{}
}

func (view View) SetTestResultMessage(msg string) {
	fmt.Println(msg)
}
