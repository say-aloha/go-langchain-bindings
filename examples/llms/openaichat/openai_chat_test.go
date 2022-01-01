
package openai

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/langchain-go/llms/openaichat"
	"log"
	"testing"
)

// To Execute EXPORT OPENAI_API_KEY=...

func TestFirstMessageChat(t *testing.T) {
	llm, err := openaichat.New()
	if err != nil {
		log.Fatal(err)
	}
	completion, err := llm.Call(context.Background(), "Hi, how are you?", []string{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}