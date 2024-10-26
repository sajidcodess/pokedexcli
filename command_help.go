package main

import "fmt"

func handleHelpCommand() error {
  fmt.Println()
  fmt.Println("Welcome to the pokedex")
  fmt.Println("Usage: ")

  for _, val := range getCommands() {
    fmt.Printf("%s: %s\n", val.name, val.description)
  }
  return nil
}
