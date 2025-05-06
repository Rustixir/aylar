
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
package main

import (
	"aylar/agent"
	"aylar/llm"
	"aylar/prompt"
	"aylar/tool"
	"fmt"
	"os"
)

func main() {
	// Initialize the language model with your API key
	openaiLLM := &llm.OpenAI{ApiKey: os.Getenv("OPENAI_API_KEY")}
	

	// Create a few-shot prompt with examples to guide the agent
	fewshot := &prompt.FewShot{
		Prefix: "Extract the following information from the job description:\n- Job Title\n- Company\n- Location\n- Experience Required\n- Salary Range\n- Application Deadline",
		Examples: []prompt.Example{
			{
				Input:  "We are seeking a Senior Software Engineer at TechCorp located in New York, NY. Candidates should have 5+ years of experience. Salary: $120,000 - $150,000. Apply by June 30, 2025.",
				Output: "Job Title: Senior Software Engineer\nCompany: TechCorp\nLocation: New York, NY\nExperience Required: 5+ years\nSalary Range: $120,000 - $150,000\nApplication Deadline: June 30, 2025",
			},
			{
				Input:  "Join HealthPlus as a Data Analyst in San Francisco, CA. Minimum 3 years experience required. Compensation between $80,000 and $100,000. Deadline to apply: July 15, 2025.",
				Output: "Job Title: Data Analyst\nCompany: HealthPlus\nLocation: San Francisco, CA\nExperience Required: 3 years\nSalary Range: $80,000 - $100,000\nApplication Deadline: July 15, 2025",
			},
		},
	}

	// Build the prompt with a new job description
	jobDescription := "Looking for a Marketing Manager at CreativeAgency based in Los Angeles, CA. Applicants should have at least 4 years of experience. Salary offered is $90,000 - $110,000. Applications close on August 10, 2025."
	customPrompt := fewshot.BuildPrompt(jobDescription)

	// Define the tools the agent can use
	tools := [
        new(tool.Search)
	]
	
	// Initialize the agent with the language model and tools
	a := agent.NewAgent(openaiLLM, tools)

	// Run the agent with the custom prompt
	answer, err := a.Run(customPrompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the agent's final answer
	fmt.Println("Extracted Information:\n", answer)
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
