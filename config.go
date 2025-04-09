package openaitranslator

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type TranslationConfig struct {
	Ctx            context.Context
	BaseURL        string
	Model          string
	SystemPrompt   string
	SourceLanguage string
	// Deprecated: This value is now deprecated in favor of max_completion_tokens,
	// and is not compatible with o1 series models.
	// refs: https://platform.openai.com/docs/api-reference/chat/create#chat-create-max_tokens
	MaxTokens           int
	MaxCompletionTokens int
	Temperature         float32
	TopP                float32
	PresencePenalty     float32
	FrequencyPenalty    float32
}

func (config *TranslationConfig) Apply(options ...Option) {
	for _, opt := range options {
		opt(config)
	}
}

const (
	DefaultModel               = openai.GPT4
	DefaultMaxCompletionTokens = 1000
	DefaultTemperature         = 0.1
	DefaultTopP                = 1.0
	DefaultPresencePenalty     = 0.0
	DefaultFrequencyPenalty    = 0.0
)

const defaultSystemPrompt = `You are a professional translator. You must follow the following rules:
1. Translate naturally and fluently, avoiding word-for-word translation
2. Do not add any explanations or notes
3. Only output the content of the translation
4. Translate the input text precisely and faithfully without adding or omitting any content`

func (config *TranslationConfig) correct() {
	if config.Ctx == nil {
		config.Ctx = context.Background()
	}
	if config.SystemPrompt == "" {
		config.SystemPrompt = defaultSystemPrompt
	}
	if config.Model == "" {
		config.Model = DefaultModel
	}
	if config.MaxCompletionTokens < 0 || config.MaxCompletionTokens > 4096 {
		config.MaxCompletionTokens = DefaultMaxCompletionTokens
	}
	if config.TopP < 0 || config.TopP > 1 {
		config.TopP = DefaultTopP
	}
	if config.Temperature < 0 || config.Temperature > 2 {
		config.Temperature = DefaultTemperature
	}
	if config.PresencePenalty < -2 || config.PresencePenalty > 2 {
		config.PresencePenalty = DefaultPresencePenalty
	}
	if config.FrequencyPenalty < -2 || config.FrequencyPenalty > 2 {
		config.FrequencyPenalty = DefaultFrequencyPenalty
	}
}

func DefaultConfig() *TranslationConfig {
	return &TranslationConfig{
		Ctx:                 context.Background(),
		SystemPrompt:        defaultSystemPrompt,
		Model:               DefaultModel,
		MaxCompletionTokens: DefaultMaxCompletionTokens,
		Temperature:         DefaultTemperature,
		TopP:                DefaultTopP,
		PresencePenalty:     DefaultPresencePenalty,
		FrequencyPenalty:    DefaultFrequencyPenalty,
	}
}
