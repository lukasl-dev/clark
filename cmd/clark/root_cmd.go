package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/lukasl-dev/clark"
	"github.com/lukasl-dev/gotilities"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "clark",
}

func init() {
	rootCmd.Run = runRoot

	set := rootCmd.Flags()
	defineFlags(set)
}

func runRoot(command *cobra.Command, args []string) {
	var err error

	if err = validateFlags(); err != nil {
		log.Fatalln(err.Error())
	}

	l := clark.NewLexer(clark.Options{
		Prefixes:         prefixes,
		PrefixIgnoreCase: prefixIgnoreCase,
		Labels:           labels,
		LabelIgnoreCase:  labelIgnoreCase,
	}, bufio.NewReader(strings.NewReader(input)))

	go l.Run()

	var out []interface{}

	for token := range l.Chan() {
		m := token.Map()

		if !advanced {
			delete(m, "code")
		}

		out = append(out, m)
	}

	var b []byte

	if prettyPrint {
		b, err = gotilities.JSONFancyMarshal(out)
	} else {
		b, err = json.Marshal(out)
	}

	if _, err = os.Stdout.Write(b); err != nil {
		log.Fatalln(err.Error())
	}
}
