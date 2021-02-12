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

// Option is a lexer configuration function.
type Option func(lexer *Lexer) error

// Prefixes defines the prefixes used by the lexer.
func Prefixes(prefixes ...string) Option {
  return func(lexer *Lexer) error {
    lexer.prefixes = prefixes
    return nil
  }
}

// Labels defines the labels used by the lexer.
func Labels(labels ...string) Option {
  return func(lexer *Lexer) error {
    lexer.labels = labels
    return nil
  }
}

// PrefixIgnoreCase defines the prefixes case sensitivity.
func PrefixIgnoreCase(b bool) Option {
  return func(lexer *Lexer) error {
    lexer.prefixIgnoreCase = b
    return nil
  }
}

// PrefixIgnoreCase defines the labels case sensitivity.
func LabelIgnoreCase(b bool) Option {
  return func(lexer *Lexer) error {
    lexer.labelIgnoreCase = b
    return nil
  }
}

// SkipPrefix defines whether the prefix should be skipped.
func SkipPrefix(b bool) Option {
  return func(lexer *Lexer) error {
    lexer.skipPrefix = b
    return nil
  }
}

// SkipLabel defines whether the label should be skipped.
func SkipLabel(b bool) Option {
  return func(lexer *Lexer) error {
    lexer.skipLabel = b
    return nil
  }
}
