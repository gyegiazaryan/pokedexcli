package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaDetails(nameOrID string) (LocationAreaDetailResponse, error) {
	endpoint := "/location-area/" + nameOrID
	url := baseURL + endpoint

	//check cache here
	data, ok := c.cache.Get(url)
	if ok {
		//cache hit
		fmt.Println("Cache hit!")
		response := LocationAreaDetailResponse{}
		err := json.Unmarshal(data, &response)
		if err != nil {
			return LocationAreaDetailResponse{}, err
		}

		return response, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaDetailResponse{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	response := LocationAreaDetailResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	//add the results to the cache
	c.cache.Add(url, data)

	return response, nil

}
