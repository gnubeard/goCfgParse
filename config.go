package main

import (
  "github.com/BurntSushi/toml"
	"fmt"
  "os"
  "log"
)

var (
  config tomlConfig
)

// This is the parent struct
type tomlConfig struct {
  Api api
}

// Example of a subortinate struct
type api struct {
  Hostname string
	Address string
  Port string
}

// You can place as many sub-structs here as you want, 
// just make sure they have a corresponding declaration in your config.toml

func loadConfig(config *tomlConfig) {
  // Config file path is stored in env for simplicity
	var cfgFilePath string
  cfgFilePath = os.Getenv("GO_TOML_CONFIG")
  if _, err := toml.DecodeFile(cfgFilePath, &config); err != nil {
    log.Println(err)
    // Configure behavior for a failed config file load
		return
  }
  log.Println(fmt.Sprintf("[Config] Successfully loaded %s", cfgFilePath))
}
