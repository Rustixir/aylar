package prompt

import (
	"aylar/tool"
	"strings"
	"testing"
)

func TestReActPromptBuilder(t *testing.T) {
	reactBuilder := &ReAct{
		Tools: []tool.Interface{
			new(tool.Search),
			new(tool.Calculator),
		},
	}

	prompt := reactBuilder.Build("What is the square root of 16?", "Initial input")

	if !strings.Contains(prompt, "Calculator, Search") {
		t.Error("Missing tool list")
	}
	if !strings.Contains(prompt, "What is the square root of 16?") {
		t.Error("Missing user question")
	}
	if !strings.Contains(prompt, "Initial input") {
		t.Error("Missing context")
	}
}
