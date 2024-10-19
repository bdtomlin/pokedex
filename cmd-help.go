package main

import (
	"fmt"
)

func cmdHelp(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.output)
	fmt.Fprintln(cfg.output, "Welcome to Pokedex")
	fmt.Fprintln(cfg.output)
	fmt.Fprintln(cfg.output, "Usage:")
	fmt.Fprintln(cfg.output)

	printSorted := func(key string, cmd cliCmd) {
		fmt.Fprintf(cfg.output, "%s: %s\n", key, cmd.description)
	}
	applyToSortedCmds(printSorted)
	return nil
}
