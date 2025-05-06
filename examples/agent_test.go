package examples

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
	tools := []tool.Interface{
		new(tool.Search),
	}

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
