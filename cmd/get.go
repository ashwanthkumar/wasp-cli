package cmd

import (
  "errors"
  "fmt"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-cli/util"
)

var GetCommand = &cobra.Command{
  Use:   "get <path>",
  Short: "Get the configuration in current path",
  Long:  `Get the configuration in current path`,
  Run: AttachHandler(performGet),
}

func init() {
    WaspCommand.AddCommand(GetCommand)
}

func performGet(args []string) error {
  if(len(args) > 1) {
    return errors.New("get takes only 1 argument")
  }
  path := args[0]
  data, err := wasp.Get(path)
  if err == nil {
    var container interface{}
    util.JsonDecode(data, &container)
    dataToString(container)
  }
  return err
}

func dataToString(data interface{}) {
  switch data.(type) {
  case []interface{}:
    for _, elem := range data.([]interface{}) {
      dataToString(elem)
    }
  case map[string]interface{}:
    fmt.Println(util.ToJson(data))
  default:
    fmt.Println(data)
  }
}
