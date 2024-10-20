package main

import (
	"errors"
	"fmt"
)

func cmdInspect(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return errors.New("A pokemon name is required")
	}
	name := args[0]

	pokemon, ok := cfg.Caught[name]
	if !ok {
		return fmt.Errorf("You don't have a pokemon named %s", name)
	}

	fmt.Fprintln(cfg.output, "Name:", pokemon.Name)
	fmt.Fprintln(cfg.output, "Height:", pokemon.Height)
	fmt.Fprintln(cfg.output, "Weight:", pokemon.Weight)
	fmt.Fprintln(cfg.output, "Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Fprintf(cfg.output, "  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Fprintln(cfg.output, "Types:")
	for _, t := range pokemon.Types {
		fmt.Fprintf(cfg.output, "  -%s\n", t.Type.Name)
	}
	return nil
}
