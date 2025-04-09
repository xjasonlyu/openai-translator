package openai_translator

import (
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type Translator struct {
	client *openai.Client
}

func NewTranslator(authToken string, opts ...TranslatorOption) *Translator {
	config := openai.DefaultConfig(authToken)
	(*TranslatorOptions)(&config).applyOptions(opts...)

	return &Translator{
		client: openai.NewClientWithConfig(config),
	}
}

type TranslatorOptions openai.ClientConfig

type TranslatorOption func(*TranslatorOptions)

func WithBaseURL(url string) TranslatorOption {
	return func(o *TranslatorOptions) {
		if url != "" {
			o.BaseURL = url
		}
	}
}

func WithHTTPClient(client *http.Client) TranslatorOption {
	return func(o *TranslatorOptions) {
		if client != nil {
			o.HTTPClient = client
		}
	}
}

func (t *TranslatorOptions) applyOptions(opts ...TranslatorOption) {
	for _, option := range opts {
		option(t)
	}
}
