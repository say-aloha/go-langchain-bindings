
package openai

import "time"

type OpenAIInput struct {
	// Model name to use
	ModelName *string // TODO: Make into Enum
	// Holds any additional parameters that are valid to pass to https://platform.openai.com/docs/api-reference/completions/create
	ModelKwargs map[string]interface{}
	// Batch size to use when passing multiple documents to generate
	BatchSize *int64
	// List of stop words to use when generating
	Stop []string