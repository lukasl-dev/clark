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

func (t Token) Map() map[string]interface{} {
	return map[string]interface{}{
		"type":  t.Type.String(),
		"code":  t.Type,
		"value": t.Value,
	}
}

func (t Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Map())
}
