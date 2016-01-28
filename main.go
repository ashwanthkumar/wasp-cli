package main

import (
  "os"
  "fmt"

  "github.com/ashwanthkumar/wasp-cli/cmd"
)

func main() {
  setupSignalHandlers()

  if err := cmd.WaspCommand.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
