package main

import (
	"testing"
)

func TestCliCommands(t *testing.T) {
	tt := []map[string]any{
		{
			"key":         "help",
			"name":        "help",
			"description": "Displays a help message",
			"cmd":         cmdHelp,
		},
		{
			"key":         "exit",
			"name":        "exit",
			"description": "Exit the Pokedex",
			"cmd":         cmdExit,
		},
		{
			"key":         "map",
			"name":        "map",
			"description": "Get a map",
			"cmd":         cmdMap,
		},
		{
			"key":         "mapb",
			"name":        "mapb",
			"description": "Get previous map",
			"cmd":         cmdMapb,
		},
	}

	cmds := cliCommands()
	for _, m := range tt {
		cmd, ok := cmds[m["key"].(string)]
		if !ok {
			t.Fatalf("missing key %s", m["key"])
		}
		if cmd.name != m["name"] {
			t.Fatalf("Expected %s, Got %s", m["name"], cmd.name)
		}
		if cmd.description != m["description"] {
			t.Fatalf("Expected %s, Got %s", m["description"], cmd.description)
		}
	}
}
