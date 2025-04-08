package openaitranslator

import (
	"fmt"
	"log"

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
	cfg.correct()
	if cfg.Url != "" {
		openaiConf.BaseURL = cfg.Url
	}
	resp, err := openai.NewClientWithConfig(openaiConf).CreateChatCompletion(cfg.Ctx, openai.ChatCompletionRequest{
		Model:            cfg.Model,
		MaxTokens:        cfg.MaxTokens,
		Temperature:      cfg.Temperature,
		TopP:             cfg.TopP,
		PresencePenalty:  cfg.PresencePenalty,
		FrequencyPenalty: cfg.FrequencyPenalty,

		Messages: generateChat(text, To, cfg),
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func generateChat(text, To string, params *TranslationConfig) []openai.ChatCompletionMessage {
	var assistantPrompt string
	To = getBaseLangCode(To)
	if name := getLangName(params.From); name == "" || name == "auto" {
		assistantPrompt = fmt.Sprintf("Please translate the following text to %s", getLangName(To))
	} else {
		assistantPrompt = fmt.Sprintf("Please translate the following text from %s to %s", name, getLangName(To))
	}
	chat := []openai.ChatCompletionMessage{
		{Role: "system", Content: params.SystemPrompt},
		{Role: "user", Content: assistantPrompt},
		{Role: "user", Content: text},
	}
	if params.Debug {
		log.Println(chat)
	}
	return chat
}
