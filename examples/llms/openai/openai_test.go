
package openai

import (
	"context"
	"fmt"
	openai "github.com/speakeasy-api/langchain-go/llms/openai"
	"log"
	"testing"
)

// To Execute EXPORT OPENAI_API_KEY=...

func TestBasicCompletion(t *testing.T) {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}