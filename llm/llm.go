package llm

type Interface interface {
	Predict(prompt string) (string, error)
}
