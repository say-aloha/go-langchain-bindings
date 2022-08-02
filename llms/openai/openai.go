package openai

import (
	"context"
	"errors"
	llms_shared "github.com/speakeasy-api/langchain-go/llms/shared"

	openai_shared "github.com/speakeasy-api/langchain-go/llms/shared/openai"
	gpt "github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/speakeasy-api/langchain-go/llms"
)

// Default Params for Open AI model
const (
	temperature      float64 = 0.7
	maxTokens        int64   = 256
	topP             float64 = 1
	frequencyPenalty float64 = 0
	presencePenalty  float64 = 0
	n                int64   = 1
	bestOf           int64   = 1
	modelName        string  = "text-davinci-003"
	batchSize        int64   = 20
	maxRetries       int     = 3
)

type OpenAI struct {
	temperature      float64
	maxTokens        int64
	topP             float64
	frequencyPenalty float64
	presencePenalty  float64
	n                int64
	bestOf           int64
	logitBias        map[string]interface{}
	streaming        bool // Streaming Unsupported Right Now
	modelName        string
	modelKwargs      map[string]interface{}
	maxRetries       int
	batchSize        int64
	stop             []string
	timeout          *time.Duration
	client           *gpt.Gpt
}

func New(args ...OpenAIInput) (*OpenAI, error) {
	if len(args) > 1 {
		return nil, errors.New("more than one config argument not supported")
	}

	input := OpenAIInput{}
	if len(args) > 0 {
		input = args[0]
	}

	openai := OpenAI{
		temperature:      temperature,
		maxTokens:        maxTokens,
		topP:             topP,
		frequencyPenalty: frequencyPenalty,
		presencePenalty:  presencePenalty,
		n:                n,
		bestOf:           bestOf,
		logitBias:        input.LogitBias,
		streaming:        input.Streaming,
		modelName:        modelName,
		modelKwargs:      input.ModelKwargs,
		batchSize:        batchSize,
		stop:             input.Stop,
		timeout:          input.Timeout,
		maxRetries:       maxRetries,
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if input.OpenAIApiKey != nil {
		apiKey = *input.OpenAIApiKey
	}

	if apiKey == "" {
		return nil, errors.New("OpenAI API key not found")
	}

	if input.ModelName != nil {
		openai.modelName = *input.ModelName
	}

	if strings.HasPrefix(openai.modelName, "gpt-3.5-turbo") || strings.HasPrefix(openai.mo