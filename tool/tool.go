package tool

type Interface interface {
	Name() string
	Description() string
	Run(input string) (string, error)
}
