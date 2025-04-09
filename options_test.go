package openai_translator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateOptions(t *testing.T) {
	ctx := context.Background()
	opts := []TranslateOption{
		WithContext(ctx),
		WithModel("model1"),
		WithSourceLanguage("lang1"),
		WithSystemPrompt("You are a translator"),
		WithMaxTokens(128), // deprecated
		WithMaxCompletionTokens(256),
		WithTemperature(0.7),
		WithTopP(0.9),
		WithPresencePenalty(0.6),
		WithFrequencyPenalty(0.4),
	}
	options := DefaultOptions()
	options.Gather(opts...)
	options.correct()

	assert.Equal(t, "model1", options.Model)
	assert.Equal(t, "lang1", options.SourceLanguage)
	assert.Equal(t, "You are a translator", options.SystemPrompt)
	assert.Equal(t, 128, options.MaxTokens)
	assert.Equal(t, 256, options.MaxCompletionTokens)
	assert.Equal(t, float32(0.7), options.Temperature)
	assert.Equal(t, float32(0.9), options.TopP)
	assert.Equal(t, float32(0.6), options.PresencePenalty)
	assert.Equal(t, float32(0.4), options.FrequencyPenalty)
	assert.Equal(t, ctx, options.Ctx)
}

func TestDefaultTranslateOptions(t *testing.T) {
	opts := []TranslateOption{
		WithModel(""),
		WithSourceLanguage(""),
		WithSystemPrompt(""),
	}
	options := DefaultOptions()
	options.Gather(opts...)
	options.correct()

	assert.Equal(t, DefaultModel, options.Model)
	assert.Equal(t, "", options.SourceLanguage)
	assert.Equal(t, defaultSystemPrompt, options.SystemPrompt)
}
