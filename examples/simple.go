package main

import (
  "bufio"
  "fmt"
  "strings"

  "github.com/lukasl-dev/clark"
)

func main() {
  input := "!help me"
  prefixes := []string{"!"}
  labels := []string{"ping", "help", "info"}

  options := clark.Options{
    Prefixes: prefixes,
    Labels:   labels,
  }

  l := clark.NewLexer(options, bufio.NewReader(strings.NewReader(input)))

  go l.Run()

  for token := range l.Chan() {
    fmt.Println(token)
  }
}
