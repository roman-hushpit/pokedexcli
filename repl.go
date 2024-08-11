package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/roman-hushpit/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	firstParameter *string
	pokedex map[string]Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if len(words) > 1 {
			parameterName := words[1]
			cfg.firstParameter = &parameterName
		}


		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore" : {
			name:        "explore",
			description: "explore the pokemons on provided location, location should be provided as a second parameter ",
			callback:    commandExplore,
		}, 
		"catch" : {
			name:        "catch",
			description: "catch by name, name should be provided as a second parameter ",
			callback:    comandCatchPokemon,
		}, 
		"inspect": {
			name:        "inspect",
			description: "inspect by name if present in pokedex, name should be provided as a second parameter ",
			callback:    commandInspect,
		},
		"pokedex" : {
			name:        "pokedex",
			description: "inspect all your pokemons",
			callback:    commandPokedex,
		}, 
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

	}
}
