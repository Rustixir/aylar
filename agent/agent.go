package agent

import (
	"aylar/llm"
	"aylar/tool"
	"fmt"
	"strings"
)

type Agent struct {
	LLM           llm.Interface
	Tools         []tool.Interface
	MaxIterations int
}

func NewAgent(llm llm.Interface, toolList []tool.Interface) *Agent {
	return &Agent{
		LLM:           llm,
		Tools:         toolList,
		MaxIterations: 5, // default cap
	}
}

func (a *Agent) Run(input string) (string, error) {
	context := input

	for i := 0; i < a.MaxIterations; i++ {
		prompt := fmt.Sprintf(`You have these tools: %s
Question: %s
What should you do?`, a.toolList(), context)

		output, err := a.LLM.Predict(prompt)
		if err != nil {
			return "", err
		}

		fmt.Printf("LLM Output: %s\n", output)

		if strings.HasPrefix(output, "Final Answer:") {
			return strings.TrimPrefix(output, "Final Answer:"), nil
		}

		handled := false
		for _, tool := range a.Tools {
			if strings.HasPrefix(output, tool.Name()) {
				input := strings.TrimPrefix(output, tool.Name()+":")
				result, err := tool.Run(strings.TrimSpace(input))
				if err != nil {
					return "", err
				}
				context += "\nObservation: " + result
				handled = true
				break
			}
		}

		if !handled {
			context += "\nObservation: Unable to handle action."
		}
	}

	return "Failed to reach final answer.", nil
}

func (a *Agent) toolList() string {
	var names []string
	for _, t := range a.Tools {
		names = append(names, t.Name())
	}
	return strings.Join(names, ", ")
}
