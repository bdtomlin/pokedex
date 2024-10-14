package main

import (
	"errors"
	"io"
	"strings"
)

type cliCmd struct {
	name        string
	description string
	callback    func(w io.Writer) error
}

func cliCommands() map[string]cliCmd {
	return map[string]cliCmd{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
	}
}

func execCommand(cmd string, w io.Writer) error {
	cmd = normalizeCmd(cmd)
	cmds := cliCommands()
	if _, ok := cmds[cmd]; !ok {
		return errors.New("invalid command")
	}
	cmds[cmd].callback(w)
	return nil
}

func normalizeCmd(cmd string) string {
	cmd = strings.ToLower(cmd)
	cmd = strings.Fields(cmd)[0]
	return cmd
}
