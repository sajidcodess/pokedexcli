package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sajidcodess/pokedexcli/internal/pokeapi"
)

type config struct {
  pokeapiClient pokeapi.Client
  nextLocationURL *string
  prevLocationURL *string
}


func startREPL(cfg *config) {
  reader := bufio.NewScanner(os.Stdin)
  for {
    fmt.Print("Pokedex -> ")
    reader.Scan() 
    words := cleanInputText(reader.Text())
    if len(words) == 0 {
      continue
    }
    args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
    command, exist := getCommands()[words[0]]
    if exist {
      err := command.callback(cfg, args...)
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
  callback func(*config, ...string) error
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
    "map": {
      name: "map",
      description: "Get the next page locations",
      callback: handleMapfCommand,
    },
    "mapb": {
      name: "mapb",
      description: "Get the previous page locations",
      callback: handleMapbCommand,
    },
    "explore": {
      name: "explore",
      description: "Explore a location",
      callback: handleExploreCommand,
    },
  }
}

