package main

import (
	"bufio"
	"fmt"
	"strings"
)

func startRepl(pd *pokedex) {
	scanner := bufio.NewScanner(pd.input)
	for {
		printPrompt(pd)
		scanner.Scan()
		execCommand(scanner.Text(), pd)
	}
}

func printPrompt(pd *pokedex) {
	fmt.Fprintln(pd.output)
	fmt.Fprint(pd.output, "pokedex > ")
}

func execCommand(cmd string, pd *pokedex) {
	cmd = normalizeCmd(cmd)
	cmds := cliCommands()
	if _, ok := cmds[cmd]; !ok {
		fmt.Fprintln(pd.output, "invalid command")
	} else {
		if err := cmds[cmd].callback(pd); err != nil {
			fmt.Fprintf(pd.output, "Error with command '%s': %s", cmd, err.Error())
		}
	}
}

func normalizeCmd(cmd string) string {
	cmd = strings.ToLower(cmd)
	cmd = strings.Fields(cmd)[0]
	return cmd
}
