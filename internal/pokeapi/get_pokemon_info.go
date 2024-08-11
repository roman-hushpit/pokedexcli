package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetPokemonInfo(pokemonName string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := PokemonResponse{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return PokemonResponse{}, err
		}

		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResponse := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResponse, nil
}
