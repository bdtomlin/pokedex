package main

import (
	"testing"
)

func TestCliCommands(t *testing.T) {
	cases := []struct {
		key         string
		name        string
		description string
		cmd         func(*config, ...string) error
	}{
		{
			key:         "cache",
			name:        "cache",
			description: "Show the cache",
			cmd:         cmdCache,
		},
		{
			key:         "catch",
			name:        "catch",
			description: "Try to catch a pokemon",
			cmd:         cmdCatch,
		},
		{
			key:         "exit",
			name:        "exit",
			description: "Exit the Pokedex",
			cmd:         cmdExit,
		},
		{
			key:         "explore",
			name:        "explore",
			description: "Explore a location",
			cmd:         cmdExit,
		},
		{
			key:         "help",
			name:        "help",
			description: "Display a help message",
			cmd:         cmdHelp,
		},
		{
			key:         "map",
			name:        "map",
			description: "Get a map",
			cmd:         cmdMap,
		},
		{
			key:         "mapb",
			name:        "mapb",
			description: "Get previous map",
			cmd:         cmdMapb,
		},
	}

	cmds := cliCommands()
	for _, c := range cases {
		cmd, ok := cmds[c.key]
		if !ok {
			t.Fatalf("missing key %s", c.key)
		}
		if cmd.name != c.name {
			t.Fatalf("Expected %s, Got %s", c.name, cmd.name)
		}
		if cmd.description != c.description {
			t.Fatalf("Expected %s, Got %s", c.description, cmd.description)
		}
	}
}
