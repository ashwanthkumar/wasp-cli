package cmd

import (
  "log"
  "fmt"
  "strings"

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
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  WaspCommand.AddCommand(LsCommand)
}

func performLS(args []string) error {
  path := strings.Join(args, " ")
  // log.Println("Using - " + wasp.Url)
  // println("Getting path from - " + path)
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
