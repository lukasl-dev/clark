/*
 * MIT License
 *
 * Copyright (c) 2021 lukas.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package clark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenizerOptions(t *testing.T) {
	actual := NewTokenizerOptions()
	expected := new(TokenizerOptions)
	assert.Equal(t, expected, actual)
}

func TestTokenizerOptions_WithPrefixes(t *testing.T) {
	var (
		prefixes = []string{"/", "-", "!"}
	)
	actual := NewTokenizerOptions().WithPrefixes(prefixes...)
	expected := &TokenizerOptions{
		prefixes: prefixes,
	}
	assert.Equal(t, expected, actual)
}

func TestTokenizerOptions_WithPrefixIgnoreCase(t *testing.T) {
	const (
		prefixIgnoreCase = true
	)
	actual := NewTokenizerOptions().WithPrefixIgnoreCase(prefixIgnoreCase)
	expected := &TokenizerOptions{
		prefixIgnoreCase: prefixIgnoreCase,
	}
	assert.Equal(t, expected, actual)
}

func TestTokenizerOptions_WithNoPrefix(t *testing.T) {
	const (
		noPrefix = true
	)
	actual := NewTokenizerOptions().WithNoPrefix(noPrefix)
	expected := &TokenizerOptions{
		noPrefix: noPrefix,
	}
	assert.Equal(t, expected, actual)
}

func TestTokenizerOptions_WithLabels(t *testing.T) {
	var (
		labels = []string{"play", "resume", "stop", "skip"}
	)
	actual := NewTokenizerOptions().WithLabels(labels...)
	expected := &TokenizerOptions{
		labels: labels,
	}
	assert.Equal(t, expected, actual)
}

func TestTokenizerOptions_WithLabelIgnoreCase(t *testing.T) {
	const (
		labelIgnoreCase = true
	)
	actual := NewTokenizerOptions().WithLabelIgnoreCase(labelIgnoreCase)
	expected := &TokenizerOptions{
		labelIgnoreCase: labelIgnoreCase,
	}
	assert.Equal(t, expected, actual)
}
