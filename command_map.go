package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	client := cfg.pokeapiClient
	result, err := client.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return fmt.Errorf("error fetching location areas: %w", err)
	}

	fmt.Println("Location Areas:")
	for _, area := range result.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = result.Next
	cfg.previousLocationAreaURL = result.Previous
	return nil
}
