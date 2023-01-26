package archive

import (
	"time"
)

type Presenter struct {
	animationChars string
	view           IView
}

func ConstructPresenter(view IView) Presenter {
	return Presenter{
		animationChars: "/-\\-",
		view:           view,
	}
}

func (presenter Presenter) PresentCompilation(receiveDone chan bool) {
	animationCharIndex := 0

	for shouldContinue := true; shouldContinue; {
		select {
		case <-receiveDone:
			shouldContinue = false
		default:
			presenter.view.updateAnimation(presenter.getAnimationFrame(animationCharIndex))
			presenter.updateAnimationCharIndex(&animationCharIndex)
		}
		presenter.waitAfterAnimationFrame()

	}
}

func (presenter Presenter) PresentCompilationSuccess() {
	msg := "\rCompiled successfully" +
		"\nRunning tests"

	presenter.view.setAnimationEndMessage(msg)
	presenter.waitAfterAnimation()
}

func (presenter Presenter) getAnimationFrame(animationCharIndex int) string {
	animationChar := presenter.animationChars[animationCharIndex]
	return "\rCompiling [" + string(animationChar) + "]"
}

func (presenter Presenter) waitAfterAnimationFrame() {
	const timeBetweenAnimationFrames = time.Millisecond * 100
	time.Sleep(timeBetweenAnimationFrames)
}

func (presenter Presenter) waitAfterAnimation() {
	const timeToWaitAfterAnimation = time.Second
	time.Sleep(timeToWaitAfterAnimation)
}
func (presenter Presenter) updateAnimationCharIndex(animationCharIndex *int) {
	const LastAnimationCharIndex = 3
	const FirstAnimationCharIndex = 0

	if *animationCharIndex == LastAnimationCharIndex {
		*animationCharIndex = FirstAnimationCharIndex
	} else {
		*animationCharIndex++
	}
}
