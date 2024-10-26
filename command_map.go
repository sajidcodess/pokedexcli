package main

import (
	"errors"
	"fmt"
)

func handleMapfCommand(cfg *config) error {
  locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
  if err != nil {
    return err
  }

  cfg.nextLocationURL = locationResp.Next
  cfg.prevLocationURL = locationResp.Previous

  for _, loc := range locationResp.Results {
    fmt.Println(loc.Name)
  }

  return nil
}

func handleMapbCommand(cfg *config) error {
  if cfg.prevLocationURL == nil {
    return errors.New("You're on the first page")
  }
  locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
  if err != nil {
    return err
  }

  cfg.nextLocationURL = locationResp.Next
  cfg.prevLocationURL = locationResp.Previous

  for _, loc := range locationResp.Results {
    fmt.Println(loc.Name)
  }

  return nil

}
