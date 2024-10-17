package main

import (
	"errors"
	"fmt"
)

func cmdMapb(cfg *config, args ...string) error {
	if cfg.Previous == "" {
		return errors.New("There is no previous map")
	}

	pMap, err := cfg.PokeApi.GetMap(cfg.Previous)
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
