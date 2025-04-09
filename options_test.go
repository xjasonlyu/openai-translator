package openai_translator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateOptions(t *testing.T) {
	ctx := context.Background()

	options := DefaultOptions()
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
	options.Gather(opts...)

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
