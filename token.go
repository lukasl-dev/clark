package clark

import (
	"encoding/json"
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

	TokenTypePrefix:   "Prefix",
	TokenTypeLabel:    "Label",
	TokenTypeText:     "Text",
	TokenTypeFlagName: "FlagName",
}

func (t TokenType) String() string {
	s, ok := tokenMap[t]

	if !ok {
		log.Panicf("missing string representation for %d", t)
	}

	return s
}

type Token struct {
	Type  TokenType `json:"type,omitempty"`
	Value string    `json:"value,omitempty"`
}

func NewToken(tokenType TokenType, value interface{}) Token {
	return Token{Type: tokenType, Value: fmt.Sprint(value)}
}

func (t Token) String() string {
	return fmt.Sprintf("%s: '%s'", t.Type.String(), t.Value)
}

func (t Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"type":  t.Type.String(),
		"code":  t.Type,
		"value": t.Value,
	})
}
