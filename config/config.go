package config

import (
  "fmt"
  "github.com/spf13/viper"
)

var c = viper.New()

func init() {
  c.SetConfigName("wasp")
  c.SetConfigType("json")
  c.AddConfigPath("$HOME/.indix/")
  err := c.ReadInConfig() // Find and read the config file

  if err != nil { // Handle errors reading the config file
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }
}

func GetWASPHost() string {
  return c.GetString("wasp.host")
}
