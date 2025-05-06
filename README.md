
# 🧠 Aylar

A lightweight Go framework for building **ReAct pattern agents** — inspired by LangChain.
It lets you combine LLMs, tools, and structured prompts to create agentic software.



## ✨ Features

✅ Implements the ReAct loop (Reason + Act + Observe)
✅ Pluggable LLM interface (`llm.LLM`)
✅ Extensible tools system (`tools.Tool`)
✅ Built-in prompt engineering system (`prompt` package)
✅ Ready-to-use tools: Calculator, Search
✅ Simple developer API for fast prototyping



## 🚀 Quick Usage

Install in your Go project:

```bash
go get github.com/Rustixir/aylar
```

Example:

```go
package examples

import (
	"aylar/agent"
	"aylar/llm"
	"aylar/tool"
	"fmt"
	"os"
)

func main() {
	openaiLLM := &llm.OpenAI{ApiKey: os.Getenv("OPENAI_API_KEY")}

	tools := []tool.Interface{
		new(tool.Calculator),
		new(tool.Search),
	}

	agent := agent.NewAgent(openaiLLM, tools)

	question := "What is 5 * 12 ? And who is the president of the USA?"
	answer, err := agent.Run(question)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Agent Final Answer:", answer)
}

```

---

## 🔧 Prompt Engineering

The `prompt` package helps you **design prompts intentionally**, not just pass raw text.

✅ **PromptTemplate** → structured templates with variables

✅ **FewShotPrompt** → include example Q\&A pairs to guide the LLM

✅ **ReActPromptBuilder** → generate well-formed ReAct prompts that explain available tools and current context


This gives you **control and consistency** over how the agent interacts with the LLM, helping it reason and choose actions effectively.

---

## 🛠 Custom Tool Example

```go
type JokeTool struct{}
func (j *JokeTool) Name() string { return "Joke" }
func (j *JokeTool) Description() string { return "Tells a joke" }
func (j *JokeTool) Run(input string) (string, error) {
    return "Why did the chicken cross the road? To get to the other side!", nil
}
```

---

## 🧪 Run Tests

```bash
go test ./...
```

---

## 📄 License

MIT License

---

If you want, I can generate the full `README.md` file for you, ready to drop into your repo — should I do that?
