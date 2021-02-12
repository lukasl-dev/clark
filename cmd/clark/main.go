package main

import (
  "log"

  "github.com/lukasl-dev/clark/cmd/clark/commands"
)

func main() {
  if err := commands.Root().Execute(); err != nil {
    log.Fatalln(err.Error())
  }
}
