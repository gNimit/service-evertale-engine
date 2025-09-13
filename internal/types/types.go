package types

// Client represents an AI model client.
type Client struct {
	Name   string `json:"name"`
	Model  string `json:"model"`
	APIKey string
	Host   string `json:"host,omitempty"`
}

// Message represents a single message in a chat conversation.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
}

// Response represents the response from the AI model.
type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Content string `json:"content"`
	Data    []byte `json:"data,omitempty"`
}

// Model represents the configuration of an AI model.
type Model struct {
	Name        string `json:"name" yaml:"name"`
	Model       string `json:"model" yaml:"model"`
	Description string `json:"description" yaml:"description"`
	Type        string `json:"type" yaml:"type"`
	Version     string `json:"version" yaml:"version"`
	BaseURL     string `json:"base_url" yaml:"base_url"`
	MaxTokens   int    `json:"max_tokens" yaml:"max_tokens"`
	RPM         int    `json:"rpm" yaml:"rpm"`
	Timeout     int    `json:"timeout" yaml:"timeout"`
}
