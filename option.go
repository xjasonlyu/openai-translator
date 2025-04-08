package openaitranslator

import (
	"context"
)

type Option func(*TranslationConfig)

func WithBaseURL(url string) Option {
	return func(tc *TranslationConfig) {
		tc.Url = url
	}
}

func WithModel(Model string) Option {
	return func(tc *TranslationConfig) {
		tc.Model = Model
	}
}

func WithSystemPrompt(prompt string) Option {
	return func(tc *TranslationConfig) {
		tc.SystemPrompt = prompt
	}
}

func WithContext(Ctx context.Context) Option {
	return func(tc *TranslationConfig) {
		tc.Ctx = Ctx
	}
}

func WithFrom(From string) Option {
	return func(tc *TranslationConfig) {
		tc.From = From
	}
}

func WithMaxTokens(MaxTokens int) Option {
	return func(tc *TranslationConfig) {
		tc.MaxTokens = MaxTokens
	}
}

func WithTemperature(Temperature float32) Option {
	return func(tc *TranslationConfig) {
		tc.Temperature = Temperature
	}
}

func WithTopP(TopP float32) Option {
	return func(tc *TranslationConfig) {
		tc.TopP = TopP
	}
}

func WithPresencePenalty(PresencePenalty float32) Option {
	return func(tc *TranslationConfig) {
		tc.PresencePenalty = PresencePenalty
	}
}

func WithFrequencyPenalty(FrequencyPenalty float32) Option {
	return func(tc *TranslationConfig) {
		tc.FrequencyPenalty = FrequencyPenalty
	}
}
