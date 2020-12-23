package main

import (
	"errors"

	"github.com/spf13/pflag"
)

var (
	ErrInputEmpty = errors.New("input cannot be empty")
)

var (
	input            string
	prefixes         []string
	prefixIgnoreCase bool
	labels           []string
	labelIgnoreCase  bool
	advanced         bool
	prettyPrint      bool
)

func defineFlags(set *pflag.FlagSet) {
	set.StringVarP(&input, "input", "i", string(readPipe(150)), "")

	/* prefix */
	set.StringSliceVarP(&prefixes, "prefix", "p", []string{}, "")
	set.BoolVar(&prefixIgnoreCase, "prefix-ignore-case", false, "")

	/* label */
	set.StringSliceVarP(&labels, "label", "l", []string{}, "")
	set.BoolVar(&labelIgnoreCase, "label-ignore-case", false, "")

	/* miscellaneous */
	set.BoolVarP(&advanced, "advanced", "a", false, "")
	set.BoolVar(&prettyPrint, "pretty-print", false, "")
}

func validateFlags() error {
	if len(input) == 0 {
		return ErrInputEmpty
	}

	return nil
}
