package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("please provide a pokemon to catch")
	}
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s... \n", pokemonName)

	client := cfg.pokeapiClient

	result, err := client.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}

	const threshold = 50
	randnNum := rand.IntN(result.BaseExperience)
	if randnNum > threshold {
		return fmt.Errorf("%s escaped", result.Name)
	}
	fmt.Printf("You caught %s ! \n", result.Name)

	cfg.Pokedex[pokemonName] = result
	fmt.Println("You may now inspect it with the inspect command.")

	return nil
}
