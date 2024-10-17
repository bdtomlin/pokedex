package main

import (
	"flag"
	"fmt"
	"os"
)

func cmdExit(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.output, "Exiting Pokedex")
	if flag.Lookup("test.v") == nil {
		os.Exit(0)
	}
	return nil
}
