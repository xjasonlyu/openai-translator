package openai_translator

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIOptions struct {
	Model        string
	SystemPrompt string
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

type TranslateOptions struct {
	Ctx            context.Context
	SourceLanguage string
	OpenAIOptions
}

func (o *TranslateOptions) Gather(options ...TranslateOption) {
	for _, option := range options {
		option(o)
	}
}

type TranslateOption func(*TranslateOptions)

func WithContext(ctx context.Context) TranslateOption {
	return func(o *TranslateOptions) {
		o.Ctx = ctx
	}
}

func WithSourceLanguage(source string) TranslateOption {
	return func(o *TranslateOptions) {
		o.SourceLanguage = source
	}
}

func WithModel(model string) TranslateOption {
	return func(o *TranslateOptions) {
		o.Model = model
	}
}

func WithSystemPrompt(prompt string) TranslateOption {
	return func(o *TranslateOptions) {
		o.SystemPrompt = prompt
	}
}

// Deprecated: use WithMaxCompletionTokens instead.
func WithMaxTokens(tokens int) TranslateOption {
	return func(o *TranslateOptions) {
		o.MaxTokens = tokens
	}
}

func WithMaxCompletionTokens(tokens int) TranslateOption {
	return func(o *TranslateOptions) {
		o.MaxCompletionTokens = tokens
	}
}

func WithTemperature(v float32) TranslateOption {
	return func(o *TranslateOptions) {
		o.Temperature = v
	}
}

func WithTopP(v float32) TranslateOption {
	return func(o *TranslateOptions) {
		o.TopP = v
	}
}

func WithPresencePenalty(v float32) TranslateOption {
	return func(o *TranslateOptions) {
		o.PresencePenalty = v
	}
}

func WithFrequencyPenalty(v float32) TranslateOption {
	return func(o *TranslateOptions) {
		o.FrequencyPenalty = v
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

func (o *TranslateOptions) correct() {
	if o.Ctx == nil {
		o.Ctx = context.Background()
	}
	if o.SystemPrompt == "" {
		o.SystemPrompt = defaultSystemPrompt
	}
	if o.Model == "" {
		o.Model = DefaultModel
	}
	if o.MaxCompletionTokens < 0 || o.MaxCompletionTokens > 4096 {
		o.MaxCompletionTokens = DefaultMaxCompletionTokens
	}
	if o.TopP < 0 || o.TopP > 1 {
		o.TopP = DefaultTopP
	}
	if o.Temperature < 0 || o.Temperature > 2 {
		o.Temperature = DefaultTemperature
	}
	if o.PresencePenalty < -2 || o.PresencePenalty > 2 {
		o.PresencePenalty = DefaultPresencePenalty
	}
	if o.FrequencyPenalty < -2 || o.FrequencyPenalty > 2 {
		o.FrequencyPenalty = DefaultFrequencyPenalty
	}
}

func DefaultOptions() *TranslateOptions {
	return &TranslateOptions{
		Ctx: context.Background(),
		OpenAIOptions: OpenAIOptions{
			Model:               DefaultModel,
			SystemPrompt:        defaultSystemPrompt,
			MaxCompletionTokens: DefaultMaxCompletionTokens,
			Temperature:         DefaultTemperature,
			TopP:                DefaultTopP,
			PresencePenalty:     DefaultPresencePenalty,
			FrequencyPenalty:    DefaultFrequencyPenalty,
		},
	}
}
