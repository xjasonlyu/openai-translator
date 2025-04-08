package openaitranslator

import (
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

	for _, unit := range []struct {
		text, from, to string
	}{
		{`Oh yeah! I'm a translator!`, "", "zh"},
		{`Oh yeah! I'm a translator!`, "", "wyw"},
		{`Oh yeah! I'm a translator!`, "", "ja"},
		{`Oh yeah! I'm a translator!`, "auto", "de"},
		{`Oh yeah! I'm a translator!`, "en", "fr"},
	} {
		result, err := Translate(
			unit.text, unit.to, apiKey,
			WithBaseURL(apiURL),
			WithSourceLanguage(unit.from),
		)
		if assert.NoError(t, err) {
			t.Log(result)
		}
	}
}
