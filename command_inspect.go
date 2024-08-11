package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error  {
	if *cfg.firstParameter == "" {
		fmt.Println("Parameter does not provided")
		return errors.New("parameter does not provided")
	}

	if pokemon, ok := cfg.pokedex[*cfg.firstParameter]; ok {
		fmt.Println(pokemon)
		return nil
	} else {
		fmt.Println("Yo dont have this pokemon")
		return nil
	}

}