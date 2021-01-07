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
