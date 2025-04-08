package openaitranslator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslationConfig(t *testing.T) {
	config := DefaultConfig()
	opts := []Option{
		WithBaseURL("url1"),
		WithModel("model1"),
		WithSourceLanguage("lang1"),
	}
	config.Apply(opts...)

	assert.Equal(t, config.BaseURL, "url1")
	assert.Equal(t, config.Model, "model1")
	assert.Equal(t, config.SourceLanguage, "lang1")
}
