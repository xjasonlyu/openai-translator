package openaitranslator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	for _, unit := range []struct {
		text, from, to string
	}{
		{`Oh yeah! I'm a translator!`, "", "zh"},
		{`Oh yeah! I'm a translator!`, "", "wyw"},
		{`Oh yeah! I'm a translator!`, "", "ja"},
		{`Oh yeah! I'm a translator!`, "", "de"},
		{`Oh yeah! I'm a translator!`, "", "fr"},
	} {
		result, err := Translate(
			unit.text, unit.to,
			os.Getenv("OPENAI_API_KEY"),
			WithFrom(unit.from),
			WithUrl(os.Getenv("OPENAI_API_URL")),
			WithModel(os.Getenv("OPENAI_MODEL")),
		)
		if assert.NoError(t, err) {
			t.Log(result)
		}
	}
}

func TestRegisterLanguage(t *testing.T) {
	t.Log(GetLangMap())
	RegisterLanguage("zh-CN", "简体中文")
	t.Log(GetLangMap())
}

var (
	_ = WithUrl
	_ = WithCtx
	_ = WithDebug
	_ = WithModel
	_ = WithFrom
	_ = WithFrequencyPenalty
	_ = WithMaxTokens
	_ = WithSystemPrompt
	_ = WithTemperature
	_ = WithPresencePenalty
	_ = WithTopP
)
