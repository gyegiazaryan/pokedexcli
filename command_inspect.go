package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a pokemon to inspect")
	}
	pokemonName := args[0]

	results, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s ", pokemonName)
	}

	fmt.Printf("Name: %s\n", results.Name)
	fmt.Printf("Height: %d\n", results.Height)
	fmt.Printf("Weight: %d\n", results.Weight)
	fmt.Printf("Stats: \n")
	for _, v := range results.Stats {
		fmt.Printf(" - %s: %v \n", v.Stat.Name, v.BaseStat)

	}
	fmt.Printf("Types: \n")
	for _, v := range results.Types {
		fmt.Printf(" - %s \n", v.Type.Name)
	}
	return nil
}
