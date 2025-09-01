package main

import (
	"pokedexcli/internal/pokeapi"
	"testing"
	"time"
)

func TestCommandInspect(t *testing.T) {
	input := "pidgey"

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		Pokedex:       make(map[string]pokeapi.PokemonResponse),
	}

	// First, catch the pokemon to ensure it is in the Pokedex
	catchErr := commandCatch(&cfg, input)
	if catchErr != nil {
		t.Errorf("expected no error catching pokemon, got %v", catchErr)
	}

	err := commandInspect(&cfg, input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}
