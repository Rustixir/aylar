package prompt

import (
	"strings"
	"testing"
)

func TestFewShotPrompt(t *testing.T) {
	fewshot := &FewShot{
		Prefix: "You are a helpful assistant.",
		Examples: []Example{
			{"What is 2+2?", "4"},
			{"What is the capital of France?", "Paris"},
		},
	}

	finalPrompt := fewshot.BuildPrompt("What is the largest planet?")

	if !strings.Contains(finalPrompt, "You are a helpful assistant.") {
		t.Error("Missing prefix in prompt")
	}
	if !strings.Contains(finalPrompt, "Q: What is 2+2?") || !strings.Contains(finalPrompt, "A: 4") {
		t.Error("Missing example 1")
	}
	if !strings.Contains(finalPrompt, "Q: What is the capital of France?") || !strings.Contains(finalPrompt, "A: Paris") {
		t.Error("Missing example 2")
	}
	if !strings.Contains(finalPrompt, "Q: What is the largest planet?") {
		t.Error("Missing user question")
	}
}
