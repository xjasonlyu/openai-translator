package openaitranslator

import (
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Translate(text, to, token string, opts ...Option) (string, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return TranslateWithConfig(text, to, token, cfg)
}

func TranslateWithConfig(text, To, token string, cfg *TranslationConfig) (string, error) {
	openaiConf := openai.DefaultConfig(token)
	cfg.init()
	if cfg.BaseURL != "" {
		openaiConf.BaseURL = cfg.BaseURL
	}
	client := openai.NewClientWithConfig(openaiConf)
	resp, err := client.CreateChatCompletion(
		cfg.Ctx,
		openai.ChatCompletionRequest{
			Model:               cfg.Model,
			MaxTokens:           cfg.MaxTokens,
			MaxCompletionTokens: cfg.MaxTokens,
			Temperature:         cfg.Temperature,
			TopP:                cfg.TopP,
			PresencePenalty:     cfg.PresencePenalty,
			FrequencyPenalty:    cfg.FrequencyPenalty,
			// translation chat messages
			Messages: generateChatMessages(text, To, cfg),
		})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("empty response choices")
	}
	return resp.Choices[0].Message.Content, nil
}

const (
	chatMessageRoleSystem = "system"
	chatMessageRoleUser   = "user"
)

func generateChatMessages(text, to string, params *TranslationConfig) []openai.ChatCompletionMessage {
	assistantPrompt := "Please translate the following text"
	if src := LookupLanguage(params.SourceLanguage); src == "" || src == "auto" {
		assistantPrompt += fmt.Sprintf(" into %s:", LookupLanguage(to))
	} else {
		assistantPrompt += fmt.Sprintf(" from %s to %s:", src, LookupLanguage(to))
	}

	messages := []openai.ChatCompletionMessage{
		{Role: chatMessageRoleSystem, Content: params.SystemPrompt},
		{Role: chatMessageRoleUser, Content: assistantPrompt},
		{Role: chatMessageRoleUser, Content: text},
	}
	return messages
}
