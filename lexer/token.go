package lexer

import (
	"fmt"
	"log"
)

type TokenType int

const (
	TokenTypeEOF TokenType = iota
	TokenTypeIllegal

	TokenTypePrefix
	TokenTypeLabel
	TokenTypeText
	TokenTypeFlagName
)

var tokenMap = map[TokenType]string{
	TokenTypeEOF:     "EOF",
	TokenTypeIllegal: "Illegal",

	TokenTypePrefix:   "TokenPrefix",
	TokenTypeLabel:    "TokenLabel",
	TokenTypeText:     "TokenText",
	TokenTypeFlagName: "TokenFlagName",
}

type Token struct {
	Type  TokenType `json:"type,omitempty"`
	Value string    `json:"value,omitempty"`
}

func NewToken(tokenType TokenType, value interface{}) Token {
	return Token{Type: tokenType, Value: fmt.Sprint(value)}
}

func (t Token) String() string {
	s, ok := tokenMap[t.Type]

	if !ok {
		log.Panicf("missing string representation for %d", t.Type)
	}

	return fmt.Sprintf("%s: '%s'", s, t.Value)
}
