package main

import (
	"errors"
	"fmt"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func cmdMapb(cfg *config) error {
	if cfg.Previous == "" {
		return errors.New("There is no previous map")
	}

	pMap, err := pokeapi.GetMap(cfg.Previous, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.Next = pMap.Next
	cfg.Previous = pMap.Previous
	for _, result := range pMap.Results {
		fmt.Fprintln(cfg.output, result.Name)
	}
	fmt.Fprintln(cfg.output, "Previous", cfg.Previous)
	fmt.Fprintln(cfg.output, "Next", cfg.Next)
	return nil
}
