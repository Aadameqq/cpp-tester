package tester

type ProgramInputProvider struct {
	inputGenerator IInputGenerator
}

func ConstructProgramInputProvider(inputGenerator InputGenerator) ProgramInputProvider {
	return ProgramInputProvider{inputGenerator: inputGenerator}
}

func (inputProvider ProgramInputProvider) Provide(index int) string {
	return inputProvider.inputGenerator.Generate(index)
}
