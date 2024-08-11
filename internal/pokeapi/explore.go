package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ExplorePokemons(areaName string) (RespPokemons, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespPokemons{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespPokemons{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemons{}, err
	}

	pokemonsResponse := RespPokemons{}
	err = json.Unmarshal(dat, &pokemonsResponse)
	if err != nil {
		return RespPokemons{}, err
	}
	c.cache.Add(url, dat)
	return pokemonsResponse, nil
}
