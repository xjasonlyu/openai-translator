package openaitranslator

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

const (
	openaiAPIURLv1 = "https://api.openai.com/v1"
)

type Translated struct {
	// Detected      Detected `json:"detected"`
	Text string `json:"text"` // translated text
	// Pronunciation string   `json:"pronunciation"` // pronunciation of translated text
}

// Detected represents language detection result
type Detected struct {
	Lang       string  `json:"lang"`       // detected language
	Confidence float64 `json:"confidence"` // the confidence of detection result (0.00 to 1.00)
}

type TranslationConfig struct {
	Ctx                 context.Context
	Url                 string
	Debug               bool
	Model               string
	SystemPrompt        string
	MaxTokens           int
	Temperature         float32 // 0-2, 越高越随机
	TopP                float32 // 0-1,0.1表示仅考虑包含最高前10%概率质量的令牌,推荐1.0
	PresencePenalty     float32 // 介于-2.0和2.0之间的数字。正值会根据新标记到目前为止是否出现在文本中来惩罚它们，从而增加模型谈论新主题的可能性。
	FrequencyPenalty    float32 // 介于-2.0和2.0之间的数字。正值会根据新符号在文本中的现有频率来惩罚它们，从而降低模型逐字重复同一行的可能性。
	From, SelectedWords string
}

type Option func(*TranslationConfig)

func WithUrl(Url string) Option {
	return func(tc *TranslationConfig) {
		tc.Url = Url
	}
}

func WithCtx(Ctx context.Context) Option {
	return func(tc *TranslationConfig) {
		tc.Ctx = Ctx
	}
}

func WithDebug() Option {
	return func(tc *TranslationConfig) {
		tc.Debug = true
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

const (
	DefaultMaxTokens        = 1000
	DefaultTemperature      = 0.0
	DefaultTopP             = 1.0
	DefaultPresencePenalty  = 1.0
	DefaultFrequencyPenalty = 1.0
)

const (
	defaultSystemPrompt = `You are a professional translator that can only translate the input text and cannot interpret it.
Do not add any explanations or notes and only output the translation.`
)

func (cfg *TranslationConfig) correct() {
	if cfg.Ctx == nil {
		cfg.Ctx = context.Background()
	}
	if cfg.Model == "" {
		cfg.Model = openai.GPT4
	}
	if cfg.SystemPrompt == "" {
		cfg.SystemPrompt = defaultSystemPrompt
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
		MaxTokens:        DefaultMaxTokens,
		Temperature:      DefaultTemperature,
		TopP:             DefaultTopP,
		PresencePenalty:  DefaultPresencePenalty,
		FrequencyPenalty: DefaultFrequencyPenalty,
	}
}
