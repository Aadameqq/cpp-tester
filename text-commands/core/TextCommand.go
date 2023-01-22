package text_commands_core

type TextCommand struct {
	name   string
	params []string
}

func ConstructTextCommand(name string, params []string) *TextCommand {
	return &TextCommand{name: name, params: params}
}

func (c TextCommand) GetName() string {
	return c.name
}

func (c TextCommand) GetParams() []string {
	return c.params
}
