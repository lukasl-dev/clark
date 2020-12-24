package main

import (
	"github.com/spf13/pflag"
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
	set.StringVarP(&input, "input", "i", string(readPipe(150)), "define raw input")

	/* prefix */
	set.StringSliceVarP(&prefixes, "prefix", "p", []string{}, "add a prefix")
	set.BoolVar(&prefixIgnoreCase, "prefix-ignore-case", false, "set prefix matching to ignore-case")

	/* label */
	set.StringSliceVarP(&labels, "label", "l", []string{}, "add a label")
	set.BoolVar(&labelIgnoreCase, "label-ignore-case", false, "set label matching to ignore-case")

	/* miscellaneous */
	set.BoolVarP(&advanced, "advanced", "a", false, "set advanced lookup")
	set.BoolVar(&prettyPrint, "pretty-print", false, "set json pretty-print")
}

func validateFlags() error {
	return nil
}
