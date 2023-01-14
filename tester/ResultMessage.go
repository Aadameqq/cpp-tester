package tester

type ResultMessage struct {
	content string
	colour  string
}

func ConstructResultMessage(content string, colour string) *ResultMessage {
	return &ResultMessage{content: content, colour: colour}
}

func (rm ResultMessage) GetContent() string {
	return rm.content
}

func (rm ResultMessage) GetColour() string {
	return rm.colour
}
