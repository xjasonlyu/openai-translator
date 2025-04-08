package openaitranslator

import "testing"

func TestRegisterLanguage(t *testing.T) {
	t.Log(GetLangMap())
	RegisterLanguage("zh-CN", "简体中文")
	t.Log(GetLangMap())
}
