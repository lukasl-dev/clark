package clark

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/lukasl-dev/gotilities"
)

type LexingFunc func(lexer *Lexer) LexingFunc

type Lexer struct {
	options   Options
	reader    *bufio.Reader
	tokens    chan Token
	lastToken *Token
}

func NewLexer(options Options, reader *bufio.Reader) *Lexer {
	return &Lexer{
		options: options,
		reader:  reader,
		tokens:  make(chan Token),
	}
}

func (l *Lexer) Chan() chan Token {
	return l.tokens
}

func (l *Lexer) Run() chan Token {
	for fn := l.lexBegin; fn != nil; {
		fn = fn(l)
	}

	defer close(l.tokens)
	return l.tokens
}

///////////////////////////////////////////////////////////////////////////
// Helper-Functions
///////////////////////////////////////////////////////////////////////////

func (l *Lexer) readAll() ([]byte, error) {
	return l.reader.Peek(l.reader.Size())
}

func (l *Lexer) submit(t TokenType, format string, v ...interface{}) {
	token := NewToken(t, fmt.Sprintf(format, v...))

	l.tokens <- token
	l.lastToken = &token
}

func (l *Lexer) error(err error) {
	if err != io.EOF {
		l.submit(TokenTypeIllegal, fmt.Sprintf("error '%s' appeared", err.Error()))
	}

	l.eof(l)
}

func (l *Lexer) discardWhitespaces() error {
	for true {
		r, _, err := l.reader.ReadRune()

		if err != nil {
			return err
		}

		if !unicode.IsSpace(r) {
			_ = l.reader.UnreadRune()
			return err
		}
	}

	return nil
}

func (l *Lexer) eof(*Lexer) LexingFunc {
	l.submit(TokenTypeEOF, "")
	return nil
}

///////////////////////////////////////////////////////////////////////////
// Lexing Methods
///////////////////////////////////////////////////////////////////////////

func (l *Lexer) lexBegin(*Lexer) LexingFunc {
	return l.lexPrefix
}

func (l *Lexer) lexPrefix(*Lexer) LexingFunc {
	if len(l.options.Prefixes) == 0 {
		return l.lexLabel
	}

	content, err := l.readAll()

	if err != nil && err != io.EOF {
		l.error(err)
		return nil
	}

	if len(content) == 0 {
		return l.eof
	}

	length := 0

	for _, prefix := range l.options.Prefixes {
		if gotilities.StringHasPrefix(string(content), prefix, l.options.PrefixIgnoreCase) && len(prefix) > length {
			length = len(prefix)
		}
	}

	b, err := ioutil.ReadAll(io.LimitReader(l.reader, int64(length)))

	if err != nil {
		l.error(err)
		return nil
	}

	l.submit(TokenTypePrefix, string(b))

	if length == len(string(content)) {
		return l.eof
	}

	return l.lexLabel
}

func (l *Lexer) lexLabel(*Lexer) LexingFunc {
	if len(l.options.Labels) == 0 {
		return l.lexFurther
	}

	content, err := l.readAll()

	if err != nil && err != io.EOF {
		l.error(err)
		return nil
	}

	if len(content) == 0 {
		return l.eof
	}

	length := 0

	for _, label := range l.options.Labels {
		if gotilities.StringHasPrefix(string(content), label, l.options.LabelIgnoreCase) && len(label) > length {
			length = len(label)
		}
	}

	b, err := ioutil.ReadAll(io.LimitReader(l.reader, int64(length)))

	if err != nil {
		l.error(err)
		return nil
	}

	l.submit(TokenTypeLabel, string(b))

	if length == len(string(content)) {
		return l.eof
	}

	return l.lexFurther
}

func (l *Lexer) lexFurther(*Lexer) LexingFunc {
	_ = l.discardWhitespaces()

	b, err := l.reader.Peek(1)

	if err != nil {
		l.error(err)
		return nil
	}

	if b[0] == '-' {
		return l.lexFlagName
	}

	return l.lexText
}

func (l *Lexer) lexText(*Lexer) LexingFunc {
	err := l.discardWhitespaces()

	if err != nil && err != io.EOF {
		l.error(err)
		return nil
	}

	s, err := l.reader.ReadString(' ')

	if len(s) == 0 {
		return l.eof
	}

	if strings.HasSuffix(s, " ") {
		s = strings.TrimRightFunc(s, func(r rune) bool {
			return unicode.IsSpace(r)
		})
	}

	if err != nil && err != io.EOF {
		l.error(err)
		return nil
	}

	l.submit(TokenTypeText, s)

	if err == io.EOF {
		return l.eof
	}

	return l.lexFurther
}

func (l *Lexer) lexFlagName(*Lexer) LexingFunc {
	name, err := l.reader.ReadString(' ')

	if err != nil && err != io.EOF {
		l.error(err)
		return nil
	}

	name = strings.TrimRightFunc(name, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	l.submit(TokenTypeFlagName, name)

	if err == io.EOF {
		return l.eof
	}

	return l.lexFurther
}
