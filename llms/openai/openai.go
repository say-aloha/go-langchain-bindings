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
	te