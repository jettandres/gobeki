package gobeki

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

// gobeki generate beki ipsum sentence
// params:
//   numWords - number of words to generate
func gobeki(numWords int) string {
	buf := bytes.Buffer{}
	for i := 0; i < numWords; i++ {
		fmt.Fprint(&buf, "hello ")
	}
	return strings.TrimSuffix(buf.String(), " ")
}

func generateRandomWord() string {
	loremsLength := int64(len(lorems))
	rand.Seed(loremsLength)
	return lorems[rand.Intn(len(lorems))]
}
