package domain

// LLMRequest represents the request structure for the LLM API.
// Documentation: https://open.bigmodel.cn/dev/api/normal-model/glm-4
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type FunctionParameters struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required,omitempty"`
}

type Tool struct {
	Type        string              `json:"type"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *FunctionParameters `json:"parameters,omitempty"`
	KnowledgeID string              `json:"knowledge_id,omitempty"`
	Prompt      string              `json:"prompt_template,omitempty"`
	Enable      *bool               `json:"enable,omitempty"`
	Query       string              `json:"search_query,omitempty"`
}

type LLMRequest struct {
	Model          string      `json:"model"`
	Messages       []Message   `json:"messages"`
	RequestID      string      `json:"request_id,omitempty"`
	DoSample       *bool       `json:"do_sample,omitempty"`
	Stream         *bool       `json:"stream,omitempty"`
	Temperature    *float64    `json:"temperature,omitempty"`
	TopP           *float64    `json:"top_p,omitempty"`
	MaxTokens      *int        `json:"max_tokens,omitempty"`
	ResponseFormat *string     `json:"response_format,omitempty"`
	Stop           []string    `json:"stop,omitempty"`
	Tools          []Tool      `json:"tools,omitempty"`
	ToolChoice     interface{} `json:"tool_choice,omitempty"`
	UserID         string      `json:"user_id,omitempty"`
}

// Option is a functional option for setting additional fields on LLMRequest.
type Option func(*LLMRequest)

// WithAPIKey sets the API key for the request.
func WithAPIKey(apiKey string) Option {
	return func(req *LLMRequest) {
		// Since API Key should be handled by the client, we do not set it here.
		// Instead, this could be used for other purposes if necessary.
	}
}

// LLMRequestBuilderService provides methods to build complex LLMRequest objects.
type LLMRequestBuilderService struct{}

// BuildLLMRequest constructs an LLMRequest object with the given parameters and options.
func (s *LLMRequestBuilderService) BuildLLMRequest(
	model string,
	messages []Message,
	options ...Option,
) LLMRequest {
	req := LLMRequest{
		Model:    model,
		Messages: messages,
	}

	// Apply optional settings through functional options
	for _, option := range options {
		option(&req)
	}

	return req
}

// Functional Options for setting optional fields in LLMRequest

// WithRequestID sets the RequestID field.
func WithRequestID(id string) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.RequestID = id
	}
}

// WithDoSample sets the DoSample field.
func WithDoSample(doSample bool) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.DoSample = &doSample
	}
}

// WithStream sets the Stream field.
func WithStream(stream bool) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.Stream = &stream
	}
}

// WithTemperature sets the Temperature field.
func WithTemperature(temperature float64) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.Temperature = &temperature
	}
}

// WithTopP sets the TopP field.
func WithTopP(topP float64) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.TopP = &topP
	}
}

// WithMaxTokens sets the MaxTokens field.
func WithMaxTokens(maxTokens int) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.MaxTokens = &maxTokens
	}
}

// WithResponseFormat sets the ResponseFormat field.
func WithResponseFormat(format string) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.ResponseFormat = &format
	}
}

// WithStop sets the Stop field.
func WithStop(stop []string) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.Stop = stop
	}
}

// WithTools sets the Tools field.
func WithTools(tools []Tool) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.Tools = tools
	}
}

// WithToolChoice sets the ToolChoice field.
func WithToolChoice(toolChoice interface{}) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.ToolChoice = toolChoice
	}
}

// WithUserID sets the UserID field.
func WithUserID(userID string) func(*LLMRequest) {
	return func(req *LLMRequest) {
		req.UserID = userID
	}
}
