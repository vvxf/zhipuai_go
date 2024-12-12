package application

import (
	"fmt"

	"github.com/vvxf/zhipuai_go/domain"
	"github.com/vvxf/zhipuai_go/infrastructure"
)

type LLMApplicationService struct {
	apiClient *infrastructure.LLMClient
}

func NewLLMApplicationService(apiURL, apiKey string) *LLMApplicationService {
	return &LLMApplicationService{
		apiClient: infrastructure.NewLLMClient(apiURL, apiKey),
	}
}

func (s *LLMApplicationService) HandleRequest(model string, messages []domain.Message) (*domain.LLMResponse, error) {
	// Prepare the API request using the builder pattern
	validateInputErr := s.validateInput(model, messages)
	if validateInputErr != nil {
		return nil, validateInputErr
	}

	request := s.buildLLMRequest(model, messages)

	// Send the request through the infrastructure layer
	response, err := s.apiClient.SendRequest(request)
	if err != nil {
		return nil, err
	}

	// Apply any post-processing or business rules on the response if needed
	if err := s.postProcessResponse(response); err != nil {
		return nil, err
	}

	return response, nil
}

// buildLLMRequest constructs an LLMRequest object using the builder service.
func (s *LLMApplicationService) buildLLMRequest(model string, messages []domain.Message) domain.LLMRequest {
	builder := &domain.LLMRequestBuilderService{}

	apiReq := builder.BuildLLMRequest(
		model,
		messages,
	)

	// apiReq := builder.BuildLLMRequest(
	// 	"glm-4-flash",
	// 	messages,
	// 	domain.WithRequestID("req_12345"),
	// 	domain.WithDoSample(true),
	// 	domain.WithStream(true),
	// 	domain.WithTemperature(0.7),
	// 	domain.WithTopP(0.9),
	// 	domain.WithMaxTokens(150),
	// 	domain.WithResponseFormat("text"),
	// 	domain.WithStop([]string{"\n"}),
	// 	domain.WithUserID("user_12345"),
	// )
	return apiReq
}

// validateInput checks if the input parameters are valid.
func (s *LLMApplicationService) validateInput(model string, messages []domain.Message) error {
	// Implement validation logic here
	if model == "" {
		return fmt.Errorf("model is required")
	}

	if len(messages) == 0 {
		return fmt.Errorf("at least one message is required")
	}
	return nil // Return nil if valid, or an error otherwise
}

// postProcessResponse applies any necessary transformations to the response.
func (s *LLMApplicationService) postProcessResponse(resp *domain.LLMResponse) error {
	// Implement post-processing logic here
	if resp == nil {
		return fmt.Errorf("response is nil")
	}
	return nil // Return nil if successful, or an error otherwise
}

// Example usage of BuildLLMRequest with functional options
// func Example() {
// 	builder := domain.LLMRequestBuilderService{}

// 	messages := []domain.Message{
// 		{Role: "user", Content: "Hello, world!"},
// 		{Role: "assistant", Content: "Hi there!"},
// 	}

// 	apiReq := builder.BuildLLMRequest(
// 		"glm-4-flash",
// 		messages,
// 		domain.WithRequestID("req_12345"),
// 		domain.WithDoSample(true),
// 		domain.WithStream(true),
// 		domain.WithTemperature(0.7),
// 		domain.WithTopP(0.9),
// 		domain.WithMaxTokens(150),
// 		domain.WithResponseFormat("text"),
// 		domain.WithStop([]string{"\n"}),
// 		domain.WithUserID("user_12345"),
// 	)

// 	// Serialize to JSON for demonstration purposes
// 	jsonData, err := json.MarshalIndent(apiReq, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error serializing:", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))
// }
