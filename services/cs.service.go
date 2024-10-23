package services

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func SendGeminiRequest(customPrompt, userMessage string) (string, error) {
	geminiAPIKey := os.Getenv("GEMINI_API")
	geminiURL := os.Getenv("GEMINI_URL")

	if geminiAPIKey == "" || geminiURL == "" {
		return "", fmt.Errorf("Gemini API key or URL is not set")
	}

	client := resty.New()

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": "Data: " + customPrompt,
					},
					{
						"text": "Answer: " + userMessage,
					},
				},
			},
		},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(geminiURL + geminiAPIKey)

	if err != nil {
		return "", fmt.Errorf("Failed to send request to Gemini API: %v", err)
	}

	var geminiResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := resty.New().JSONUnmarshal(resp.Body(), &geminiResponse); err != nil {
		return "", fmt.Errorf("Failed to parse Gemini API response: %v", err)
	}

	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		return geminiResponse.Candidates[0].Content.Parts[0].Text, nil
	}

	return "No text found in Gemini API response", nil
}

