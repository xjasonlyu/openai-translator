package openaitranslator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupLanguage(t *testing.T) {
	tests := []struct {
		code, lang string
	}{
		{"en", "English"},
		{"en-US", "American English"},
		{"en-CA", "Canadian English"},
		{"en-GB", "British English"},
		{"zh", "Chinese"},
		{"zh-Hans", "Simplified Chinese"},
		{"zh-Hant", "Traditional Chinese"},
		{"zh-CN", "Chinese (China)"},
		{"zh-TW", "Chinese (Taiwan)"},
		{"yue", "Cantonese"},
		// fallback languages
		{"unknown", "unknown"},
		// registered languages
		{"wyw", "中文（古文-文言文）"},
	}
	for _, tt := range tests {
		assert.Equal(t, LookupLanguage(tt.code), tt.lang)
	}
}
