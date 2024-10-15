package main

import (
	"io"
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
