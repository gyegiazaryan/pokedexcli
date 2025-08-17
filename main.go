package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{}

func main() {

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
	}

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		//cleaned := cleanInput(input)

		switch input {
		case "exit":
			commands["exit"].callback()
		case "help":
			commands["help"].callback()
		}
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, val := range commands {
		fmt.Printf("%v: %v \n", val.name, val.description)
	}

	return nil
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	x := strings.Fields(strings.ToLower(trimmed))
	return x
}
