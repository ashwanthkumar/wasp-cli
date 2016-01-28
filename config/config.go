package config

import (
  "log"
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
  notNullConfiguration("wasp.host")
  return c.GetString("wasp.host")
}

func GetAuthToken() string {
  notNullConfiguration("wasp.token")
  return c.GetString("wasp.token")
}

func notNullConfiguration(key string) {
  if c.GetString(key) == "" {
    log.Fatal(key + " configuration is not found in " + c.ConfigFileUsed())
  }
}
