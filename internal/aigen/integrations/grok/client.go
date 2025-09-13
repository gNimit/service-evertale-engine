package grok

import (
	"github.com/gNimit/service-evertale-engine/internal/types"
)

// ClientGrok is a wrapper around the aigen.Client for Grok-specific functionality.
type ClientGrok struct {
	Client *types.Client
}

// IGrokClient defines the interface for interacting with the Grok AI model.
type IGrokClient interface {
	GenerateText(prompt string) (string, error)
}

// NewGrokClient initializes and returns a new Grok AI model client.
