package main

import (
  "fmt"
  "log"
)

func main() {
  // Load the configuration file into memory
  loadConfig(&config)

  // Display configration
  log.Println(fmt.Sprintf("Hostname: %s", config.Api.Hostname))
  log.Println(fmt.Sprintf("Address: %s", config.Api.Address))
  log.Println(fmt.Sprintf("Port: %s", config.Api.Port))
}
