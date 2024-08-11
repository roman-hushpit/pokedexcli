package main

import (
	"fmt"
	"errors"
	"math/rand"
		"strings"
)

type Pokemon struct {
	Name string
	Height int
	Weight int
	Stats []Stat
	Types []string
}

func (p Pokemon) String() string {
	var stats strings.Builder
	for _, stat := range p.Stats {
		stats.WriteString(fmt.Sprintf("  -%s: %d\n", strings.ToLower(stat.Name), stat.Value))
	}

	var types strings.Builder
	for _, typ := range p.Types {
		types.WriteString(fmt.Sprintf("  - %s\n", typ))
	}

	return fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n%sTypes:\n%s", 
		p.Name, p.Height, p.Weight, stats.String(), types.String())
}

type Stat struct {
	Name string
	Value int
}


func comandCatchPokemon(cfg *config) error {
	if *cfg.firstParameter == "" {
		fmt.Println("Parameter does not provided")
		return errors.New("parameter does not provided")
	}

	if _, ok := cfg.pokedex[*cfg.firstParameter]; ok {
		fmt.Printf("%s already caught\n", *cfg.firstParameter)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.firstParameter)

	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(*cfg.firstParameter)
	if err != nil {
		return err
	}

	expirience := pokemonInfo.BaseExperience
	catchChance := rand.Intn(expirience) 
	if catchChance > expirience / 2 {
		stats :=make([]Stat, 0, 10)
		for _, s := range pokemonInfo.Stats {
			stats = append(stats, Stat{
				Name: s.Stat.Name,
				Value: s.BaseStat,
			})
		}

		types :=make([]string, 0, 10)
		for _, t := range pokemonInfo.Types {
			types = append(types, t.Type.Name)
		}

		cfg.pokedex[*cfg.firstParameter] = Pokemon{
			Name: pokemonInfo.Name,
			Height: pokemonInfo.Height,
			Weight: pokemonInfo.Weight,
			Stats: stats,
			Types: types,
		}
		fmt.Printf("%s was caught!\n", *cfg.firstParameter)
		return nil
	} else {
		fmt.Printf("%s escaped!\n", *cfg.firstParameter)
		return nil
	}
}