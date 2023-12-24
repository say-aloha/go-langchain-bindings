
package openaichat

import (
	"context"
	"errors"
	openai_shared "github.com/speakeasy-api/langchain-go/llms/shared/openai"
	gpt "github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"math"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/speakeasy-api/langchain-go/llms"
)

// Default Params for Open AI model
const (
	temperature      float64 = 1
	topP             float64 = 1
	frequencyPenalty float64 = 0
	presencePenalty  float64 = 0
	n                int64   = 1
	modelName        string  = "gpt-3.5-turbo"
	maxRetries       int     = 3
)

type OpenAIChat struct {
	temperature      float64
	maxTokens        int64
	topP             float64
	frequencyPenalty float64
	presencePenalty  float64
	n                int64
	logitBias        map[string]interface{}
	streaming        bool // Streaming Unsupported Right Now
	modelName        string
	modelKwargs      map[string]interface{}
	maxRetries       int
	stop             []string
	prefixMessages   []ChatMessage
	timeout          *time.Duration
	client           *gpt.Gpt
}

func New(args ...OpenAIChatInput) (*OpenAIChat, error) {
	if len(args) > 1 {