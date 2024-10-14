package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func cmdExit(w io.Writer) error {
	fmt.Fprintln(w, "Exiting Pokedex")
	if flag.Lookup("test.v") == nil {
		os.Exit(0)
	}
	return nil
}
