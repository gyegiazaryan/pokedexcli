package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(nameOrID string) (PokemonResponse, error) {
	endpoint := "/pokemon/" + nameOrID

	url := baseURL + endpoint

	//check cache here
	data, ok := c.cache.Get(url)
	if ok {
		//cache hit
		fmt.Println("Cache hit!")
		response := PokemonResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return PokemonResponse{}, err
		}

		return response, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	response := PokemonResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return PokemonResponse{}, err
	}

	//add the results to the cache
	c.cache.Add(url, data)

	return response, nil

}
