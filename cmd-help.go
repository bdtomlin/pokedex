package main

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

func cmdHelp(w io.Writer) error {
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Welcome to Pokedex")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w)

	cmdSl := [][]string{}
	for _, cmd := range cliCommands() {
		cmdSl = append(cmdSl, []string{cmd.name, cmd.description})
	}

	slices.SortFunc(cmdSl, func(a, b []string) int {
		return strings.Compare(strings.ToLower(a[0]), strings.ToLower(b[0]))
	})
	for _, cmd := range cmdSl {
		fmt.Fprintf(w, "%s: %s\n", cmd[0], cmd[1])
	}
	return nil
}
