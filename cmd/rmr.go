package cmd

import (
  "log"
  "errors"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-cli/config"
  "github.com/ashwanthkumar/wasp-cli/util"
)

var RmrCommand = &cobra.Command{
  Use:   "rmr <path>",
  Short: "Delete a config key (recursively)",
  Long:  `Delete a config key (recursively)`,
  Run: AttachHandler(performRmr),
}

func init() {
  WaspCommand.AddCommand(RmrCommand)
}

var yes bool
func prepareFlagsForRmr() {
  RmrCommand.PersistentFlags().BoolVarP(
    &rawData, "yes", "y", false, "Don't prompt just delete it")
}

func performRmr(args []string) error {
  wasp.SetToken(config.GetAuthToken())
  if(len(args) != 1) {
    return errors.New("rmr takes exactly 1 argument")
  }
  path := args[0]
  if path == "" {
    log.Fatal("Whoever you are, you just can't delete the entire DB. Not happening man! I'm sorry.")
  }
  var err error
  if yes || util.AskBool("Do you want to delete everything under " + path + "? [yN]: ") {
    log.Println("Deleting the configuration recursively at " + path)
    _, err = wasp.Delete(path)
  }
  return err
}
