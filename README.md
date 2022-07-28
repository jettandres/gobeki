# Gobeki

A simple TDD exercise with Go

Requirements: Go v18+

## Preview

[![asciicast](https://asciinema.org/a/511350.svg)](https://asciinema.org/a/511350)

## Install

```sh
go get github.com/jettandres/gobeki
```

## Usage

```go
import (
  gobeki "github.com/jettandres/gobeki"
  "fmt"
)

func main() {
  wordCount := 20
  paragraphCount := 1

  lorems := gobeki.Gobeki(wordCount, paragraphCount)

  fmt.Println(lorems)
}
```
