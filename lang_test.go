package openai_translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupLanguage(t *testing.T) {
	tests := []struct {
		code, lang string
	}{
		{"en", "English"},
		{"eng", "English"},
		{"fra", "French"},
		{"en-US", "American English"},
		{"en-CA", "Canadian English"},
		{"en-GB", "British English"},
		{"ja", "Japanese"},
		{"JPN", "Japanese"},
		{"zh", "Chinese"},
		{"zh-Hans", "Simplified Chinese"},
		{"zh-Hant", "Traditional Chinese"},
		{"zh-CN", "Chinese (China)"},
		{"zh_cn", "Chinese (China)"},
		{"zh-TW", "Chinese (Taiwan)"},
		{"yue", "Cantonese"},
		// fallback languages
		{"unknown", "unknown"},
		// registered languages
		{"wyw", "中文（古文-文言文）"},
		{"chs", "Simplified Chinese"},
		{"CHT", "Traditional Chinese"},
	}
	for _, tt := range tests {
		assert.Equal(t, LookupLanguage(tt.code), tt.lang)
	}
}

func TestRegisterLanguage(t *testing.T) {
	tests := []struct {
		code, lang string
	}{
		{"code1", "Code 1"},
		{"code2", "Code 2"},
		{"code3", "Code 3"},
	}
	for _, tt := range tests {
		assert.Equal(t, LookupLanguage(tt.code), tt.code)
	}
	for _, tt := range tests {
		RegisterLanguage(tt.code, tt.lang)
	}
	for _, tt := range tests {
		assert.Equal(t, LookupLanguage(tt.code), tt.lang)
	}
}
