// zhipuai_go/api/api.go
package api

import (
	"github.com/vvxf/zhipuai_go/application"
	"github.com/vvxf/zhipuai_go/domain"
)

type Message = domain.Message

func NewLLMApplicationService(apiURL, apiKey string) *application.LLMApplicationService {
	return application.NewLLMApplicationService(apiURL, apiKey)
}
