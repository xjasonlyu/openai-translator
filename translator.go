package openai_translator

import (
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type Translator struct {
	client *openai.Client
	//config *openai.ClientConfig
}

func NewTranslator(authToken string, opts ...TranslatorOption) *Translator {
	config := openai.DefaultConfig(authToken)
	for _, opt := range opts {
		opt(&config)
	}

	return &Translator{
		client: openai.NewClientWithConfig(config),
	}
}

type TranslatorOption func(*openai.ClientConfig)

func WithBaseURL(url string) TranslatorOption {
	return func(c *openai.ClientConfig) {
		if url != "" {
			c.BaseURL = url
		}
	}
}

func WithHTTPClient(client *http.Client) TranslatorOption {
	return func(c *openai.ClientConfig) {
		if client != nil {
			c.HTTPClient = client
		}
	}
}
