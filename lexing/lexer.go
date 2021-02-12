/*
 * Copyright 2021 lukasl-dev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lexing

import (
  "bufio"
  "fmt"
  "io"
  "strings"
  "unicode"

  "github.com/lukasl-dev/clark/lexing/token"
)

type Reader func(l *Lexer) Reader

// Lexer is a lexicographical reader.
type Lexer struct {
  // reader holds the content to read.
  reader *bufio.Reader

  // tokens is the channel where the tokens are sent in.
  tokens chan token.Token

  // lastToken is the last token sent to the channel.
  lastToken *token.Token

  // prefixes defines the available prefixes.
  prefixes []string

  // labels defines the available labels.
  labels []string

  // prefixIgnoreCase indicates whether the prefixes are not case sensitive.
  prefixIgnoreCase bool

  // labelIgnoreCase indicates whether the labels are not case sensitive.
  labelIgnoreCase bool

  // skipPrefix defines whether the prefix should be skipped.
  skipPrefix bool

  // skipLabel defines whether the prefix should be skipped.
  skipLabel bool
}

// NewLexer returns a new Lexer.
func NewLexer(reader io.Reader, options ...Option) (*Lexer, error) {
  l := &Lexer{}
  l.Reset(reader)

  for _, option := range options {
    if err := option(l); err != nil {
      return nil, err
    }
  }

  return l, nil
}

// Present returns true if the reader is not nil.
func (l *Lexer) Present() bool {
  return l.reader != nil
}

// Reset resets the lexer to be reading from reader.
func (l *Lexer) Reset(reader io.Reader) {
  l.tokens = make(chan token.Token)
  l.reader = bufio.NewReader(reader)
  l.lastToken = nil
}

// Chan returns the token channel in which the tokens get streamed in.
func (l *Lexer) Chan() chan token.Token {
  return l.tokens
}

// accept sends a new token into the token channel.
func (l *Lexer) accept(t token.Type, v ...interface{}) {
  var val string
  if len(v) == 1 {
    val = fmt.Sprintf("%s", v[0])
  }
  in := token.Token{Type: t, Val: val}
  l.tokens <- in
  l.lastToken = &in
}

// error sends an error into token channel.
// As soon as io.EOF is passed, then eof gets executed.
func (l *Lexer) error(err error) {
  if err == io.EOF {
    l.accept(token.EOF)
  } else {
    l.accept(token.UnexpectedEOF)
  }
}

// curr returns the content of the reader.
func (l *Lexer) curr() ([]byte, error) {
  return l.reader.Peek(l.reader.Size())
}

func (l *Lexer) hasPrefix(s string, ignoreCase bool) bool {
  curr, _ := l.curr()
  if ignoreCase {
    return strings.HasPrefix(strings.ToLower(string(curr)), strings.ToLower(s))
  }
  return strings.HasPrefix(string(curr), s)
}

// whitespace returns true, if the next rune is a whitespace.
func (l *Lexer) whitespace() (bool, error) {
  r, _, err := l.reader.ReadRune()
  _ = l.reader.UnreadRune()
  return unicode.IsSpace(r), err
}

// discardWhitespaces reads the reader's whitespaces and returns the fallback Reader.
// If max <= 0, all available whitespaces get discarded.
// The read runes get put in the dest rune-slice.
func (l *Lexer) discardWhitespaces(dest []rune, max int, fallback Reader) Reader {
  return func(lexer *Lexer) Reader {
    for i := 1; i <= max || max <= 0; i++ {
      r, _, err := lexer.reader.ReadRune()

      if !unicode.IsSpace(r) {
        _ = lexer.reader.UnreadRune()
        break
      }

      if err != nil {
        if err == io.EOF {
          dest = append(dest, r)
        } else {
          _ = lexer.reader.UnreadRune()
          lexer.error(err)
          return nil
        }
      }
    }
    return fallback
  }
}

// Lex reads the passed reader lexicographically and returns the lexed tokens.
func (l *Lexer) Lex() chan token.Token {
  if !l.Present() {
    panic("the lexer must be present to start lexing")
  }

  defer close(l.tokens)

  for reader := l.begin(); reader != nil; {
    reader = reader(l)
  }

  return l.tokens
}
