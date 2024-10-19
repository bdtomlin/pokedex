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
		return err
	}

	fmt.Fprintf(cfg.output, "Throwing a Pokeball at %s...\n", pokemon.Name)

	if catchIsSuccessful(cfg, pokemon.BaseExperience) {
		cfg.Caught[pokemon.Name] = pokemon
		fmt.Fprintf(cfg.output, "%s was caught!\n", pokemon.Name)
	} else {
		fmt.Fprintf(cfg.output, "%s escaped!\n", pokemon.Name)
	}
	return nil
}

func catchIsSuccessful(cfg *config, baseExperience int) bool {
	if cfg.testing {
		return cfg.testingRandSuccess
	}
	factor := baseExperience/25 + 1
	randNum := rand.IntN(factor)
	return randNum == 0
}
