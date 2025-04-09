package openai_translator

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	var (
		apiKey = os.Getenv("OPENAI_API_KEY")
		apiURL = os.Getenv("OPENAI_API_URL")
	)
	if apiKey == "" {
		t.SkipNow()
	}

	translator := NewTranslator(
		apiKey,
		WithBaseURL(apiURL),
		WithHTTPClient(http.DefaultClient),
	)

	for _, unit := range []struct {
		text, from, to string
	}{
		{`Oh yeah! I'm a translator!`, "", "zh"},
		{`Oh yeah! I'm a translator!`, "", "ZH-TW"},
		{`Oh yeah! I'm a translator!`, "", "wyw"},
		{`Oh yeah! I'm a translator!`, "", "ja"},
		{`Oh yeah! I'm a translator!`, "auto", "de"},
		{`Oh yeah! I'm a translator!`, "en", "fr"},
	} {
		result, err := translator.TranslateText(
			unit.text, unit.to,
			WithSourceLanguage(unit.from),
		)
		if assert.NoError(t, err) {
			t.Log(result)
		}
	}
}
