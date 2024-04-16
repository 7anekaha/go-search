package lib

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune)bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func Analize(text string) []string {
	str := Tokenize(text)
	str = LowercaseFilter(str)
	str = StopWordsFilter(str)
	str = StemmerFilter(str)
	return str
}