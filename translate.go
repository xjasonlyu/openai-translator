package openaitranslator

import (
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Translate(text, to, token string, options ...Option) (string, error) {
	config := DefaultConfig()
	config.Apply(options...)
	return TranslateWithConfig(text, to, token, config)
}

func TranslateWithConfig(text, to, token string, config *TranslationConfig) (string, error) {
	clientConfig := openai.DefaultConfig(token)
	if config.BaseURL != "" {
		clientConfig.BaseURL = config.BaseURL
	}
	config.correct()
	client := openai.NewClientWithConfig(clientConfig)
	resp, err := client.CreateChatCompletion(
		config.Ctx,
		openai.ChatCompletionRequest{
			Model:               config.Model,
			MaxTokens:           config.MaxTokens,
			MaxCompletionTokens: config.MaxCompletionTokens,
			Temperature:         config.Temperature,
			TopP:                config.TopP,
			PresencePenalty:     config.PresencePenalty,
			FrequencyPenalty:    config.FrequencyPenalty,
			// translation chat messages
			Messages: generateChatMessages(text, to, config),
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

func generateChatMessages(text, to string, config *TranslationConfig) []openai.ChatCompletionMessage {
	assistantPrompt := "Please translate the following text"
	if src := LookupLanguage(config.SourceLanguage); src == "" || src == "auto" {
		assistantPrompt += fmt.Sprintf(" into %s:", LookupLanguage(to))
	} else {
		assistantPrompt += fmt.Sprintf(" from %s to %s:", src, LookupLanguage(to))
	}

	messages := []openai.ChatCompletionMessage{
		{Role: chatMessageRoleSystem, Content: config.SystemPrompt},
		{Role: chatMessageRoleUser, Content: assistantPrompt},
		{Role: chatMessageRoleUser, Content: text},
	}
	return messages
}
