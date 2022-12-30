package compiler

import "fmt"

type View struct {
}

func ConstructView() View {
	return View{}
}

func (view View) updateAnimation(msg string) {
	fmt.Print(msg)
}

func (view View) setAnimationEndMessage(msg string) {
	fmt.Println(msg)
}
