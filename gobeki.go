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
func gobeki(numWords int) string {
	buf := bytes.Buffer{}
	for i := 0; i < numWords; i++ {
		fmt.Fprintf(&buf, "%s ", generateRandomWord())
	}
	return strings.TrimSuffix(buf.String(), " ")
}

func generateRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return lorems[rand.Intn(len(lorems))]
}
