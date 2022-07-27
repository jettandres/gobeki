package gobeki

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// gobeki generate beki ipsum sentence
// params:
//   numWords - number of words to generate
func gobeki(numWords int, paragraphs int) string {
	buf := bytes.Buffer{}

	for p := 0; p < paragraphs; p++ {
		for i := 0; i < numWords; i++ {
			fmt.Fprintf(&buf, "%s ", generateRandomWord())
		}

		fmt.Fprintln(&buf)
	}

	formattedResult := strings.TrimSuffix(buf.String(), "\n")
	formattedResult = strings.TrimSuffix(formattedResult, " ")
	return formattedResult
}

func generateRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return lorems[rand.Intn(len(lorems))]
}
