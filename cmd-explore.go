package main

import (
	"errors"
	"fmt"
)

func cmdExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("A location is required to explore")
	}
	name := args[0]

	la, err := cfg.PokeApi.GetLocation(name)
	if err != nil {
		return err
	}

	for _, encounter := range la.PokemonEncounters {
		pokemon := encounter.Pokemon
		fmt.Fprintln(cfg.output, "-", pokemon.Name)
	}
	return nil
}
