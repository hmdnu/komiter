package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hmdnu/komiter/internal/env"
)

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func HttpRequest(prompt string) (string, error) {
	client := &http.Client{}
	payloadStruct := map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	payload, err := json.Marshal(payloadStruct)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, env.API_URL, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("X-goog-api-key", env.API_TOKEN)
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("generating commit message...")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	parsedResponse, err := parseGeminiResponse(res)
	if err != nil {
		return "", err
	}
	return parsedResponse.Candidates[0].Content.Parts[0].Text, nil
}

func parseGeminiResponse(res *http.Response) (*GeminiResponse, error) {
	var geminiRes GeminiResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gemini API error: status %d, body: %s", res.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &geminiRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Gemini response: %w\nbody: %s", err, string(body))
	}
	if len(geminiRes.Candidates) == 0 {
		return nil, fmt.Errorf("no candidates in Gemini response: %s", string(body))
	}
	if len(geminiRes.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no parts in first candidate: %s", string(body))
	}
	return &geminiRes, nil
}
