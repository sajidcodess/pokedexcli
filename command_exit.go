package main

import "os"

func handleExitCommand(cfg *config, args ...string) error {
  os.Exit(0)
  return nil
}
