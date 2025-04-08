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

func (cfg *TranslationConfig) init() {
	if cfg.Ctx == nil {
		cfg.Ctx = context.Background()
	}
	if cfg.SystemPrompt == "" {
		cfg.SystemPrompt = defaultSystemPrompt
	}
	if cfg.Model == "" {
		cfg.Model = DefaultModel
	}
	if cfg.MaxTokens < 0 || cfg.MaxTokens > 4096 {
		cfg.MaxTokens = DefaultMaxTokens
	}
	if cfg.TopP < 0 || cfg.TopP > 1 {
		cfg.TopP = DefaultTopP
	}
	if cfg.Temperature < 0 || cfg.Temperature > 2 {
		cfg.Temperature = DefaultTemperature
	}
	if cfg.PresencePenalty < -2 || cfg.PresencePenalty > 2 {
		cfg.PresencePenalty = DefaultPresencePenalty
	}
	if cfg.FrequencyPenalty < -2 || cfg.FrequencyPenalty > 2 {
		cfg.FrequencyPenalty = DefaultFrequencyPenalty
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
