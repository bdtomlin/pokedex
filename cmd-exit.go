package main

import (
	"flag"
	"fmt"
	"os"
)

func cmdExit(pd *pokedex) error {
	fmt.Fprintln(pd.output, "Exiting Pokedex")
	if flag.Lookup("test.v") == nil {
		os.Exit(0)
	}
	return nil
}
