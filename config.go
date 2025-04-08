package openaitranslator

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type TranslationConfig struct {
	Ctx                 context.Context
	Url                 string
	Model               string
	SystemPrompt        string
	MaxTokens           int
	Temperature         float32 // 0-2, 越高越随机
	TopP                float32 // 0-1,0.1表示仅考虑包含最高前10%概率质量的令牌,推荐1.0
	PresencePenalty     float32 // 介于-2.0和2.0之间的数字。正值会根据新标记到目前为止是否出现在文本中来惩罚它们，从而增加模型谈论新主题的可能性。
	FrequencyPenalty    float32 // 介于-2.0和2.0之间的数字。正值会根据新符号在文本中的现有频率来惩罚它们，从而降低模型逐字重复同一行的可能性。
	From, SelectedWords string
}

const (
	DefaultModel            = openai.GPT4
	DefaultMaxTokens        = 1000
	DefaultTemperature      = 0.0
	DefaultTopP             = 1.0
	DefaultPresencePenalty  = 1.0
	DefaultFrequencyPenalty = 1.0
)

const (
	defaultSystemPrompt = `You are a professional translator. You must follow the following rules:
1. Translate naturally and fluently, avoiding word-for-word translation
2. Do not add any explanations or notes
3. Only output the content of the translation`
)

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
