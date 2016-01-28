package cmd

import (
  "fmt"
  "errors"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-go/util"
)

var LsCommand = &cobra.Command{
  Use:   "ls <path>",
  Short: "List keys in the current path",
  Long:  `List keys in the current path`,
  Run: AttachHandler(performLS),
}

func init() {
  WaspCommand.AddCommand(LsCommand)
}

func performLS(args []string) error {
  if(len(args) > 1) {
    return errors.New("ls takes only 1 argument")
  }
  path := args[0]
  data, err := wasp.List(path)
  if err == nil {
    var keys []string
    util.JsonDecode(data, &keys)

    for _, key := range keys {
      fmt.Println(key)
    }
  }
  return err
}
