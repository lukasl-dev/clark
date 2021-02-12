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
