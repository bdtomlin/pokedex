package main

import (
	"fmt"
	"slices"
	"strings"
)

func cmdHelp(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.output)
	fmt.Fprintln(cfg.output, "Welcome to Pokedex")
	fmt.Fprintln(cfg.output)
	fmt.Fprintln(cfg.output, "Usage:")
	fmt.Fprintln(cfg.output)

	cmdSl := [][]string{}
	for _, cmd := range cliCommands() {
		cmdSl = append(cmdSl, []string{cmd.name, cmd.description})
	}

	slices.SortFunc(cmdSl, func(a, b []string) int {
		return strings.Compare(strings.ToLower(a[0]), strings.ToLower(b[0]))
	})
	for _, cmd := range cmdSl {
		fmt.Fprintf(cfg.output, "%s: %s\n", cmd[0], cmd[1])
	}
	return nil
}
