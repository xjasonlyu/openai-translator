package openaitranslator

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type TranslationConfig struct {
	Ctx              context.Context
	BaseURL          string
	Model            string
	SystemPrompt     string
	SourceLanguage   string
	MaxTokens        int
	Temperature      float32
	TopP             float32
	PresencePenalty  float32
	FrequencyPenalty float32
}

func (config *TranslationConfig) Apply(options ...Option) {
	for _, opt := range options {
		opt(config)
	}
}

const (
	DefaultModel            = openai.GPT4
	DefaultMaxTokens        = 1000
	DefaultTemperature      = 0.1
	DefaultTopP             = 1.0
	DefaultPresencePenalty  = 0.0
	DefaultFrequencyPenalty = 0.0
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
	if config.MaxTokens < 0 || config.MaxTokens > 4096 {
		config.MaxTokens = DefaultMaxTokens
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
		Ctx:              context.Background(),
		SystemPrompt:     defaultSystemPrompt,
		Model:            DefaultModel,
		MaxTokens:        DefaultMaxTokens,
		Temperature:      DefaultTemperature,
		TopP:             DefaultTopP,
		PresencePenalty:  DefaultPresencePenalty,
		FrequencyPenalty: DefaultFrequencyPenalty,
	}
}
