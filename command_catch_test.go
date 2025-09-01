package main

import (
	"pokedexcli/internal/pokeapi"
	"testing"
	"time"
)

func TestCommandCatch(t *testing.T) {
	input := "abra"

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		Pokedex:       make(map[string]pokeapi.PokemonResponse),
	}

	err := commandCatch(&cfg, input)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}
