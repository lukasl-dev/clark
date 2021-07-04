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

type TokenizerOptions struct {
	prefixes         []string
	prefixIgnoreCase bool
	noPrefix         bool
	labels           []string
	labelIgnoreCase  bool
	noLabel          bool
}

func NewTokenizerOptions() *TokenizerOptions {
	return &TokenizerOptions{}
}

func minimizeTokenizerOptions(opts []*TokenizerOptions) *TokenizerOptions {
	if len(opts) > 0 {
		return opts[0]
	}
	return NewTokenizerOptions()
}

func (opts *TokenizerOptions) WithPrefixes(prefixes ...string) *TokenizerOptions {
	opts.prefixes = prefixes
	return opts
}

func (opts *TokenizerOptions) WithPrefixIgnoreCase(prefixIgnoreCase bool) *TokenizerOptions {
	opts.prefixIgnoreCase = prefixIgnoreCase
	return opts
}

func (opts *TokenizerOptions) WithNoPrefix(noPrefix bool) *TokenizerOptions {
	opts.noPrefix = noPrefix
	return opts
}

func (opts *TokenizerOptions) WithLabels(labels ...string) *TokenizerOptions {
	opts.labels = labels
	return opts
}

func (opts *TokenizerOptions) WithLabelIgnoreCase(labelIgnoreCase bool) *TokenizerOptions {
	opts.labelIgnoreCase = labelIgnoreCase
	return opts
}

func (opts *TokenizerOptions) WithNoLabel(noLabel bool) *TokenizerOptions {
	opts.noLabel = noLabel
	return opts
}
