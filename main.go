package main

import (
	"time"

	"github.com/sajidcodess/pokedexcli/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.NewClient(5*time.Second, time.Second*5)
	cfg := &config{
		pokeapiClient: pokeclient,
	}
	startREPL(cfg)
}
