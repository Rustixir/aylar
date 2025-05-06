package prompt

import "testing"

func TestTemplate_Format(t *testing.T) {
	tpl := &Template{
		Template: "Hello, {{name}}! What is {{task}}?",
	}
	filled := tpl.Format(map[string]string{
		"name": "Alice",
		"task": "your favorite color",
	})
	if filled != "Hello, Alice! What is your favorite color?" {
		t.Error("Template format failed")
	}
}
