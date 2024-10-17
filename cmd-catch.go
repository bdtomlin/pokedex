package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func cmdCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("A pokemon name is required")
	}
	name := args[0]

	pokemon, err := cfg.PokeApi.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("Error retrieving pokemon: %w", err)
	}

	factor := pokemon.BaseExperience/25 + 1

	fmt.Fprintf(cfg.output, "Throwing a Pokeball at %s...\n", pokemon.Name)

	randNum := rand.IntN(factor)
	if randNum == 0 {
		cfg.Caught[pokemon.Name] = pokemon
		fmt.Fprintf(cfg.output, "%s was caught!", pokemon.Name)
	} else {
		fmt.Fprintf(cfg.output, "%s escaped!", pokemon.Name)
	}
	return nil
}
