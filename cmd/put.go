package cmd

import (
  "errors"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-cli/util"
  "github.com/ashwanthkumar/wasp-cli/config"
)

var PutCommand = &cobra.Command{
  Use:   "put <path> <value>",
  Short: "Put a value against a path",
  Long:  `Put a value against a path`,
  Run: AttachHandler(performPut),
}

func init() {
  prepareFlags()
  WaspCommand.AddCommand(PutCommand)
}

var rawData bool
var stdIn bool
func prepareFlags() {
  PutCommand.PersistentFlags().BoolVarP(
    &rawData, "raw", "r", false, "Put the value as it without parsing it as JSON")
  PutCommand.PersistentFlags().BoolVarP(
    &stdIn, "stdin", "", false, "Read from STDIN instead of value from command line")
}

func performPut(args []string) error {
  wasp.SetToken(config.GetAuthToken())
  if(len(args) != 2 && !stdIn) {
    return errors.New("put takes exactly 2 arguments")
  }
  path := args[0]
  value := ""
  if stdIn {
    value = util.ReadFullyFromStdin()
  } else {
    value = args[1]
  }

  if !rawData {
    value = util.ToJson(value)
  }
  _, err := wasp.Put(path, value)
  return err
}
