package llms

type LLMResult struct {
	Generations [][]Generation
	LLMOutput   map[string]i