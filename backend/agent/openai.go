package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

// Communication with OpenAI
type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Format the payload
type openaiRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

// Response
type openaiResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func getPlanFromGPT(prompt string) (string, error) {
	reqBody := openaiRequest{
		Model: "gpt-4",
		Messages: []chatMessage{
			{Role: "system", Content: "You're a software planner. Break user prompts into step-by-step dev tasks."},
			{Role: "user", Content: prompt},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to encode OpenAI request: %v", err)
	}

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("OpenAI request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("OpenAI error: %s", string(body))
	}

	var gptResp openaiResponse
	err = json.NewDecoder(res.Body).Decode(&gptResp)
	if err != nil {
		return "", fmt.Errorf("failed to decode OpenAI response: %v", err)
	}

	return gptResp.Choices[0].Message.Content, nil
}
