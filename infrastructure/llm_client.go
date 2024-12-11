package infrastructure

import (
	"bytes"
	"encoding/json"
	"net/http"
	"zhipuai_go/domain"
)

type LLMClient struct {
	APIURL     string
	APIKey     string // Add API key field
	httpClient *http.Client
}

func NewLLMClient(apiURL, apiKey string) *LLMClient {
	return &LLMClient{
		APIURL:     apiURL,
		APIKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// Helper function to send API request with API Key in Authorization header
func (c *LLMClient) SendRequest(request domain.LLMRequest) (*domain.LLMResponse, error) {
	// Marshal the request into JSON
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", c.APIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set headers including the Authorization header with API Key
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal the response
	var llmResp domain.LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		return nil, err
	}

	return &llmResp, nil
}
