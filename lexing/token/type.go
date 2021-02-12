/*
 *    Copyright 2021 lukasl-dev
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package token

import "log"

const (
	EOF Type = iota
	UnexpectedEOF
	TypePrefix
	TypeLabel
	TypeArgument
	TypeFlag
)

var typeMap = map[Type]string{
	EOF:           "EOF",
	UnexpectedEOF: "UnexpectedEOF",
	TypePrefix:    "Prefix",
	TypeLabel:     "Label",
	TypeArgument:  "Argument",
	TypeFlag:      "Flag",
}

type Type int

// String returns the string representation of t.
func (t Type) String() string {
	s, ok := typeMap[t]
	if !ok {
		log.Panicf("missing string representation for %d", t)
	}
	return s
}

// EOF returns true, if the type is EOF or a subtype of EOF.
func (t Type) EOF() bool {
	return t == EOF || t == UnexpectedEOF
}
