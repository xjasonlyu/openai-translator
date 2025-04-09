package openai_translator

import (
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (t *Translator) TranslateText(text, to string, opts ...TranslateOption) (string, error) {
	options := DefaultOptions()
	options.Gather(opts...)
	options.correct()

	resp, err := t.client.CreateChatCompletion(
		options.Ctx,
		openai.ChatCompletionRequest{
			Model:               options.Model,
			MaxTokens:           options.MaxTokens,
			MaxCompletionTokens: options.MaxCompletionTokens,
			Temperature:         options.Temperature,
			TopP:                options.TopP,
			PresencePenalty:     options.PresencePenalty,
			FrequencyPenalty:    options.FrequencyPenalty,
			// translation messages
			Messages: generateChatMessages(text, to, options),
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

func generateChatMessages(text, to string, options *TranslateOptions) []openai.ChatCompletionMessage {
	assistantPrompt := "Please translate the following text"
	if src := LookupLanguage(options.SourceLanguage); src == "" || src == "auto" {
		assistantPrompt += fmt.Sprintf(" into %s:", LookupLanguage(to))
	} else {
		assistantPrompt += fmt.Sprintf(" from %s to %s:", src, LookupLanguage(to))
	}

	messages := []openai.ChatCompletionMessage{
		{Role: chatMessageRoleSystem, Content: options.SystemPrompt},
		{Role: chatMessageRoleUser, Content: assistantPrompt},
		{Role: chatMessageRoleUser, Content: text},
	}
	return messages
}
