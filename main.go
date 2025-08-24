package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

var commands = map[string]cliCommand{}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Returns the help options",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "prints a list of map locations",
			callback:    commandMap},
		"mapb": {
			name:        "mapb",
			description: "prints the previous list of map locations",
			callback:    commandMapb,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "exit":
			commands["exit"].callback(&cfg)
		case "help":
			commands["help"].callback(&cfg)
		case "map":
			commands["map"].callback(&cfg)
		case "mapb":
			commands["mapb"].callback(&cfg)
		}

	}

}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, val := range commands {
		fmt.Printf("%v: %v \n", val.name, val.description)
	}

	return nil
}

func commandMap(cfg *config) error {
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

func commandMapb(cfg *config) error {
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
func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	x := strings.Fields(strings.ToLower(trimmed))
	return x
}
