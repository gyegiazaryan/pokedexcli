package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a location area to explore")
	}
	locationAreaName := args[0]
	//TODO: make this a user input
	//location := "canalave-city-area"
	client := cfg.pokeapiClient

	result, err := client.GetLocationAreaDetails(locationAreaName)
	if err != nil {
		return fmt.Errorf("error fetching location areas: %w", err)
	}

	fmt.Printf("Exploring %s\n", result.Name)
	fmt.Println("Found Pokemon:")
	for _, v := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	return nil
}
