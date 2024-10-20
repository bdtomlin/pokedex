package main

import (
	"errors"
	"fmt"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func cmdCatch(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return errors.New("A pokemon name is required")
	}
	name := args[0]

	pokemon, err := cfg.PokeApi.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Fprintf(cfg.output, "Throwing a Pokeball at %s...\n", pokemon.Name)

	maybeCatch(cfg, pokemon)
	return nil
}

func maybeCatch(cfg *config, p pokeapi.Pokemon) {
	factor := p.BaseExperience/25 + 2
	randNum := cfg.RandIntFunc(factor)
	if randNum == 1 {
		capture(cfg, p)
		fmt.Fprintf(cfg.output, "%s was caught!\n", p.Name)
	} else {
		fmt.Fprintf(cfg.output, "%s escaped!\n", p.Name)
	}
}

func capture(cfg *config, p pokeapi.Pokemon) {
	cfg.Caught[p.Name] = p
}
