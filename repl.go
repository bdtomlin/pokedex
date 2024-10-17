package main

import (
	"bufio"
	"fmt"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(cfg.input)
	for {
		printPrompt(cfg)
		scanner.Scan()
		execCommand(scanner.Text(), cfg)
	}
}

func printPrompt(cfg *config) {
	fmt.Fprintln(cfg.output)
	fmt.Fprint(cfg.output, "pokedex > ")
}

func execCommand(cmd string, cfg *config) {
	cmd, args := normalizeCmd(cmd)
	if cmd == "" {
		fmt.Fprintln(cfg.output)
	}
	cmds := cliCommands()
	if _, ok := cmds[cmd]; !ok {
		fmt.Fprintln(cfg.output, "invalid command")
	} else {
		if err := cmds[cmd].callback(cfg, args...); err != nil {
			fmt.Fprintf(cfg.output, "Error with command '%s': %s", cmd, err.Error())
		}
	}
}

func normalizeCmd(cmd string) (string, []string) {
	cmd = strings.ToLower(cmd)
	split := strings.Fields(cmd)
	var args []string
	if len(split) > 0 {
		cmd = split[0]
		args = split[1:]
	}
	return cmd, args
}
