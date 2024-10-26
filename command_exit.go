package main

import "os"

func handleExitCommand() error {
  os.Exit(0)
  return nil
}
