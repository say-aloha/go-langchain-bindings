package openai

import (
	"fmt"
	"net/http"
)

type OpenAIError struct {
	error

	statusCode int
	status     strin