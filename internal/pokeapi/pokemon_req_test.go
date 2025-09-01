package pokeapi

import (
	"testing"
	"time"
)

func TestGetPokemon(t *testing.T) {

	client := NewClient(10 * time.Second)

	_, err := client.GetPokemon("pikachu")
	if err != nil {
		t.Errorf("error fetching pokemon: %v", err)
	}

	_, err = client.GetPokemon("1")
	if err != nil {
		t.Errorf("error fetching pokemon by id: %v", err)
	}

	_, err = client.GetPokemon("notapokemon")
	if err == nil {
		t.Errorf("expected error fetching invalid pokemon, got none")
	}
}
