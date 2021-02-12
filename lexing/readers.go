package lexing

import (
  "io"
  "io/ioutil"
  "strings"
  "unicode"

  "github.com/lukasl-dev/clark/lexing/token"
)

// begin starts the lexing progress.
func (l *Lexer) begin() Reader {
  return readPrefix()
}

// readPrefix reads the prefix lexicographically until io.EOF.
// A prefix is defined as the (first) component that introduces the command-line.
// The most common prefix for chat systems like discord would be '!'.
func readPrefix() Reader {
  return func(l *Lexer) Reader {
    if l.skipPrefix {
      return readLabel()
    }

    whitespace, err := l.whitespace()

    if err != nil {
      l.error(ErrNoLabel)
      return nil
    }

    if whitespace {
      return l.discardWhitespaces(nil, 0, readArgumentsAndFlags())
    }

    size := -1
    for _, prefix := range l.prefixes {
      if l.hasPrefix(prefix, l.prefixIgnoreCase) && len(prefix) > size {
        size = len(prefix)
      }
    }

    if size == -1 {
      l.error(ErrNoPrefix)
      return nil
    }

    prefix, err := ioutil.ReadAll(io.LimitReader(l.reader, int64(size)))

    if err != nil {
      l.error(err)
      return nil
    }

    l.accept(token.TypePrefix, string(prefix))

    return readLabel()
  }
}

// readLabel reads the label lexicographically until io.EOF.
// A label is defined as the (second) component that identifies the command.
// Examples: 'help', 'configure command'
func readLabel() Reader {
  return func(l *Lexer) Reader {
    if l.skipLabel {
      return readArgumentsAndFlags()
    }

    size := -1
    for _, label := range l.labels {
      if l.hasPrefix(label, l.labelIgnoreCase) && len(label) > size {
        size = len(label)
      }
    }

    if size == -1 {
      l.error(ErrNoLabel)
      return nil
    }

    label, err := ioutil.ReadAll(io.LimitReader(l.reader, int64(size)))

    if err != nil {
      l.error(err)
      return nil
    }

    l.accept(token.TypeLabel, string(label))

    return readArgumentsAndFlags()
  }
}

// readArgumentsAndFlags reads all arguments and flags lexicographically until io.EOF.
func readArgumentsAndFlags() Reader {
  return func(l *Lexer) Reader {
    whitespace, err := l.whitespace()

    if err != nil {
      l.error(err)
      return nil
    }

    if whitespace {
      return l.discardWhitespaces(nil, 0, readArgumentsAndFlags())
    }

    var builder strings.Builder
    var quoted bool
    var flag bool

    for i := 0; true; i++ {
      var r rune
      r, _, err = l.reader.ReadRune()

      if i == 0 && r == '-' {
        flag = true
      }

      if !flag && (r == '\'' || r == '"') {
        if quoted {
          break
        }
        quoted = true
        continue
      }

      if (unicode.IsSpace(r) && (flag || !quoted)) || err != nil {
        break
      }

      builder.WriteRune(r)
    }

    t := token.TypeArgument
    if flag {
      t = token.TypeFlag
    }

    l.accept(t, builder.String())

    if err != nil {
      l.error(err)
      return nil
    }

    return readArgumentsAndFlags()
  }
}
