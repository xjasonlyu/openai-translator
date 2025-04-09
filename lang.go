package openai_translator

import (
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var _languages = sync.Map{}

func RegisterLanguage(code, lang string) {
	_languages.Store(code, lang)
}

func LookupLanguage(code string) string {
	if code == "" {
		return ""
	}
	if lang, ok := _languages.Load(code); ok {
		return lang.(string)
	}
	tag, err := language.Parse(code)
	if err != nil {
		return code // fallback to original language code.
	}
	eng := language.English
	return display.Tags(eng).Name(tag)
}

func init() {
	RegisterLanguage("wyw", "中文（古文-文言文）")
	RegisterLanguage("chs", "Simplified Chinese")
	RegisterLanguage("cht", "Traditional Chinese")
}
