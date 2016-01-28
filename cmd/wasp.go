package cmd

import (
  "log"
  "os"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-go/config"
  "github.com/ashwanthkumar/wasp-go/client"
)

var wasp = client.WASP {
  Url: config.GetWASPHost(),
  AuthToken: config.GetAuthToken(),
}

// Main command for Cobra.
var WaspCommand = &cobra.Command{
  Use:   "wasp <args>",
  Short: "Command line client to WASP",
  Long:  `Command line client to WASP`,
}

type CommandHandler func(args []string) error

func AttachHandler(handler CommandHandler) func (*cobra.Command, []string) {
  return func (cmd *cobra.Command, args []string) {
    err := handler(args)
    if err != nil {
      log.Printf("[Error] %s", err.Error())
      cmd.Help()
      os.Exit(1)
    }
  }
}

func init() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
