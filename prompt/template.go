package prompt

import "strings"

type Template struct {
	Template string
}

func (t *Template) Format(vars map[string]string) string {
	result := t.Template
	for k, v := range vars {
		placeholder := "{{" + vars[k] + "}}"
		result = strings.ReplaceAll(result, "{{"+placeholder+"}}", v)
	}
	return result
}
