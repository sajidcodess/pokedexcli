package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startREPL() {
  reader := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex -> ")
    reader.Scan() 
    words := cleanInputText(reader.Text())
    if len(words) == 0 {
      continue
    }
    command, exist := getCommands()[words[0]]
    if exist {
      err := command.callback()
      if err != nil {
        fmt.Println(err)
      }
      continue
    } else {
      fmt.Println("Unknown Command")
    }

  }
}

func cleanInputText(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output)
  return words
}

type CliCommand struct {
  name string
  description string
  callback func() error
}

func getCommands() map[string]CliCommand {
  return map[string]CliCommand {
    "help": {
      name: "help",
      description: "display help here",
      callback: handleHelpCommand,
    },
    "exit": {
      name: "exit",
      description: "Exit the program",
      callback: handleExitCommand,
    },
  }
}

