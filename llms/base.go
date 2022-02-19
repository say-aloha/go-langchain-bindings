package llms

import "context"

type LLM interface {
	Generate(ctx context.Context, prompts [