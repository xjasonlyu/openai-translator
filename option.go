package openaitranslator

import (
	"context"
)

type Option func(*TranslationConfig)

func WithContext(ctx context.Context) Option {
	return func(config *TranslationConfig) {
		config.Ctx = ctx
	}
}

func WithBaseURL(url string) Option {
	return func(config *TranslationConfig) {
		config.BaseURL = url
	}
}

func WithModel(model string) Option {
	return func(config *TranslationConfig) {
		config.Model = model
	}
}

func WithSystemPrompt(prompt string) Option {
	return func(config *TranslationConfig) {
		config.SystemPrompt = prompt
	}
}

func WithSourceLanguage(source string) Option {
	return func(config *TranslationConfig) {
		config.SourceLanguage = source
	}
}

func WithMaxTokens(tokens int) Option {
	return func(config *TranslationConfig) {
		config.MaxTokens = tokens
	}
}

func WithTemperature(v float32) Option {
	return func(config *TranslationConfig) {
		config.Temperature = v
	}
}

func WithTopP(v float32) Option {
	return func(config *TranslationConfig) {
		config.TopP = v
	}
}

func WithPresencePenalty(v float32) Option {
	return func(config *TranslationConfig) {
		config.PresencePenalty = v
	}
}

func WithFrequencyPenalty(v float32) Option {
	return func(config *TranslationConfig) {
		config.FrequencyPenalty = v
	}
}
