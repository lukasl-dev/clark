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

package commands

import (
  "encoding/json"
  "log"
  "strings"
  "time"

  "github.com/lukasl-dev/clark"
  "github.com/lukasl-dev/clark/cmd/clark/pipe"
  "github.com/lukasl-dev/clark/lexing"
  "github.com/lukasl-dev/clark/lexing/token"
  "github.com/spf13/cobra"
)

func Root() *cobra.Command {
  var (
    input            string
    prefixes         []string
    labels           []string
    prefixIgnoreCase bool
    labelIgnoreCase  bool
    skipPrefix       bool
    skipLabel        bool
    prettyPrint      bool
  )

  c := &cobra.Command{
    Use: "root",
    Run: func(cmd *cobra.Command, args []string) {
      l, err := clark.NewLexer(
        strings.NewReader(input),
        lexing.Prefixes(prefixes...),
        lexing.Labels(labels...),
        lexing.PrefixIgnoreCase(prefixIgnoreCase),
        lexing.LabelIgnoreCase(labelIgnoreCase),
        lexing.SkipPrefix(skipPrefix),
        lexing.SkipLabel(skipLabel),
      )

      if err != nil {
        log.Fatalln(err.Error())
      }

      go l.Lex()

      var tokens []token.Token
      for t := range l.Chan() {
        tokens = append(tokens, t)
      }

      var b []byte

      if prettyPrint {
        b, err = json.MarshalIndent(tokens, "", "  ")
      } else {
        b, err = json.Marshal(tokens)
      }

      if err != nil {
        log.Fatalln(err.Error())
      }

      _ = pipe.Write(b)
    },
  }

  p, err := pipe.Read(time.Duration(1) * time.Millisecond)

  if err != nil {
    log.Fatalln(err.Error())
  }

  c.Flags().StringVarP(&input, "input", "i", string(p), "")
  c.Flags().StringSliceVarP(&prefixes, "prefix", "p", nil, "")
  c.Flags().StringSliceVarP(&labels, "label", "l", nil, "")
  c.Flags().BoolVar(&labelIgnoreCase, "label-ignore-case", false, "")
  c.Flags().BoolVar(&prefixIgnoreCase, "prefix-ignore-case", false, "")
  c.Flags().BoolVar(&skipLabel, "skip-label", false, "")
  c.Flags().BoolVar(&skipPrefix, "skip-prefix", false, "")
  c.Flags().BoolVar(&prettyPrint, "pretty-print", false, "")

  return c
}
