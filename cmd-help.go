package main

import (
	"fmt"
	"slices"
	"strings"
)

func cmdHelp(pd *pokedex) error {
	fmt.Fprintln(pd.output)
	fmt.Fprintln(pd.output, "Welcome to Pokedex")
	fmt.Fprintln(pd.output)
	fmt.Fprintln(pd.output, "Usage:")
	fmt.Fprintln(pd.output)

	cmdSl := [][]string{}
	for _, cmd := range cliCommands() {
		cmdSl = append(cmdSl, []string{cmd.name, cmd.description})
	}

	slices.SortFunc(cmdSl, func(a, b []string) int {
		return strings.Compare(strings.ToLower(a[0]), strings.ToLower(b[0]))
	})
	for _, cmd := range cmdSl {
		fmt.Fprintf(pd.output, "%s: %s\n", cmd[0], cmd[1])
	}
	return nil
}
