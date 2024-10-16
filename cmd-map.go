package main

import (
	"fmt"
	"time"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func cmdMap(cfg *config) error {
	url := cfg.Url
	if cfg.Next != "" {
		url = cfg.Next
	}

	t0 := time.Now()
	pMap, err := pokeapi.GetMap(url, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.Next = pMap.Next
	cfg.Previous = pMap.Previous
	for _, result := range pMap.Results {
		fmt.Fprintln(cfg.output, result.Name)
	}
	fmt.Fprintln(cfg.output)
	fmt.Fprintln(cfg.output, "Previous", cfg.Previous)
	fmt.Fprintln(cfg.output, "Next", cfg.Next)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	return nil
}
