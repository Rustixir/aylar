package prompt

import "strings"

type Example struct {
	Input  string
	Output string
}

type FewShot struct {
	Prefix   string
	Examples []Example
}

func (f *FewShot) BuildPrompt(input string) string {
	var sb strings.Builder
	sb.WriteString(f.Prefix + "\n\n")
	for _, ex := range f.Examples {
		sb.WriteString("Q: " + ex.Input + "\n")
		sb.WriteString("A: " + ex.Output + "\n\n")
	}
	sb.WriteString("Q: " + input + "\n")
	return sb.String()
}
