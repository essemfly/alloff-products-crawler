package utils

import (
	"testing"

	"golang.org/x/text/language"
)

func TestTranslateText(t *testing.T) {
	TranslateText(language.Korean.String(), "My name is socks")
}
