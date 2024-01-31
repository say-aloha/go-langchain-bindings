package openai

import (
	"fmt"
	"net/http"
)

type OpenAIError struct {
	error

	statusCode int
	status     string
}

func (e *OpenAIError) Error() string {
	return fmt.Sprintf("error in call to o