package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	Pokedex                 map[string]pokeapi.PokemonResponse
}

func startRepl() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		Pokedex:       make(map[string]pokeapi.PokemonResponse),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		commands := getCommands()
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(&cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "prints the previous list of map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "explore a location area in detail",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
