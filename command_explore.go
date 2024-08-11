package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if *cfg.firstParameter == "" {
		fmt.Println("Parameter does not provided")
		return errors.New("parameter does not provided")
	}
	fmt.Println("Exploring pastoria-city-area...")

	exploreResponse, err := cfg.pokeapiClient.ExplorePokemons(*cfg.firstParameter)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range exploreResponse.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}