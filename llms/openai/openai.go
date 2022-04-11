package openai

import (
	"context"
	"errors"
	llms_shared "github.com/speakeasy-api/langchain-go/llms/shared"

	openai_shared "github.com/speakeasy-api/langchain-go/llms/shared/openai"
	gpt "github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"math"
	"ne