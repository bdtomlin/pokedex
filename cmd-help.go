package main

import (
	"fmt"
	"io"
)

func cmdHelp(w io.Writer) error {
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Welcome to Pokedex")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w)

	for _, cmd := range cliCommands() {
		fmt.Fprintf(w, "%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
