package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMapb(cfg *config, args ...string) error {
	client := cfg.pokeapiClient

	if cfg.previousLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	result, err := client.GetLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas:")
	for _, area := range result.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = result.Next
	cfg.previousLocationAreaURL = result.Previous
	return nil
}
