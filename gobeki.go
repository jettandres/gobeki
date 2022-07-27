package gobeki

import (
	"bytes"
	"fmt"
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
