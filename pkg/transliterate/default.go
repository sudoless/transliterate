package transliterate

import (
	transliterateData "go.sdls.io/transliterate/pkg/transliterate-data"
	transliterateLang "go.sdls.io/transliterate/pkg/transliterate-lang"
)

var defaultReplacer = Replacer{
	Lang: transliterateLang.Data,
	Data: transliterateData.Data,
}

// Transliterate is a helper function around a default Replacer using the transliterate_data.Data and transliterate_lang.Data.
func Transliterate(text, lang string) string {
	return defaultReplacer.Transliterate(text, lang)
}
