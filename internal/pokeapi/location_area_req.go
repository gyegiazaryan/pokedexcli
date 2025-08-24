package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	url := baseURL + endpoint

	if pageURL != nil {
		url = *pageURL
	}

	//check cache here
	data, ok := c.cache.Get(url)
	if ok {
		//cache hit
		fmt.Println("Cache hit!")
		response := LocationAreaResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return response, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	response := LocationAreaResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	//add the results to the cache
	c.cache.Add(url, data)

	return response, nil

}
