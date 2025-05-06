package prompt

import (
	"aylar/tool"
	"fmt"
	"strings"
)

type ReAct struct {
	Tools []tool.Interface
}

func (r *ReAct) Build(userInput string, context string) string {
	var toolInfos []string
	for _, tool := range r.Tools {
		info := fmt.Sprintf("%s: %s", tool.Name(), tool.Description())
		toolInfos = append(toolInfos, info)
	}

	toolList := strings.Join(toolInfos, "\n")

	return fmt.Sprintf(`You can use the following tools:
%s

Question: %s
Context so far: %s

Decide your next action.`, toolList, userInput, context)
}
