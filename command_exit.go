package main

import "os"

func handleExitCommand(cfg *config) error {
  os.Exit(0)
  return nil
}
