package main

import (
	"fmt"
)

func cmdPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("This command doesn't accept additional arguments")
	}
	if len(cfg.Caught) == 0 {
		return fmt.Errorf("Your Pokedex is empty!")
	}
	for _, pokemon := range cfg.Caught {
		fmt.Fprintln(cfg.output, "  -", pokemon.Name)
	}
	return nil
}
