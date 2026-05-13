package ai

import "context"

// LLMProvider define el contrato estricto Multi-Cloud para interactuar con IAs.
type LLMProvider interface {
	// Ask sends a prompt to the LLM and returns the generated text response.
	Ask(ctx context.Context, prompt string) (string, error)

	// Name returns the provider's human-readable name (e.g., "Google-Gemini").
	Name() string
}
