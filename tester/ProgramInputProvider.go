package tester

type ProgramInputProvider struct {
	inputGenerator IInputGenerator
}

func (inputProvider ProgramInputProvider) Provide(index int) string {
	return inputProvider.inputGenerator.Generate(index)
}
