package gobeki

import (
	"strings"
	"testing"
)

func TestGenerateBeki(t *testing.T) {
	t.Run("generate sentence with n number words", func(t *testing.T) {
		wordCount := 2
		got := gobeki(wordCount)

		length := len(strings.Split(got, " "))

		if length != wordCount {
			t.Errorf("got %d, want %d", length, wordCount)
		}
	})
}
