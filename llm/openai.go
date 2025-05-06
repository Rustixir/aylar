package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenAI struct {
	ApiKey string
}

type openAIRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type openAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func (o *OpenAI) Predict(prompt string) (string, error) {
	reqBody := openAIRequest{
		Model:     "text-davinci-003",
		Prompt:    prompt,
		MaxTokens: 150,
	}

	bodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	var apiResp openAIResponse
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return "", err
	}

	if len(apiResp.Choices) > 0 {
		return apiResp.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}
