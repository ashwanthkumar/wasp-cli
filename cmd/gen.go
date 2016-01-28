package cmd

import (
  "fmt"
  "os"
  "errors"
  "text/template"

  "github.com/spf13/cobra"
  "github.com/ashwanthkumar/wasp-cli/util"
)

var GenCommand = &cobra.Command{
  Use:   "gen <path> <tmpl_file>",
  Short: "Generate files using Go templates from WASP configs",
  Long:  `Generate files using Go templates from WASP configs`,
  Run: AttachHandler(performGen),
}

func init() {
  WaspCommand.AddCommand(GenCommand)
}

func performGen(args []string) (err error) {
  if(len(args) != 2) {
    return errors.New("gen takes exactly 2 arguments")
  }

  path := args[0]
  tmplFile := args[1]
  data, err := wasp.Get(path)
  if err == nil {
    var container interface{}
    util.JsonDecode(data, &container)

    // catch the template.Must's panic and return that as an error
    defer func() {
      if recovered := recover(); recovered != nil { //catch
        err = errors.New(fmt.Sprintf("Exception: %v\n", recovered))
      }
    }()
    return template.Must(template.ParseFiles(tmplFile)).Execute(os.Stdout, container)
  }

  return err
}
