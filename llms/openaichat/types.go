
package openaichat

import (
	"github.com/speakeasy-api/langchain-go/llms/openai"
)

type OpenAIChatInput struct {
	// ChatGPT messages to pass as a prefix to the prompt
	PrefixMessages []ChatMessage

	openai.OpenAIInput
}

type ChatMessage struct {
	Content string
	Role    ChatMessageRoleEnum
}