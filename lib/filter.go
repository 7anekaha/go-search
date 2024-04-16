package lib

import (
	"strings"

	stemES "github.com/kljensen/snowball/spanish"
)

func LowercaseFilter(text []string) []string {
	str := make([]string, len(text))
	for i, s := range text {
		str[i] = strings.ToLower(s)
	}
	return str
}

func StopWordsFilter(text []string) []string {
	str := make([]string, len(text))
	for _, s := range text {
		if !stemES.IsStopWord(s) {
			str = append(str, s)
		}
	}
	return str
}

func StemmerFilter(text []string) []string {
	str := make([]string, len(text))
	for i, s := range text {
		str[i] = stemES.Stem(s, false)
	}
	return str
}
