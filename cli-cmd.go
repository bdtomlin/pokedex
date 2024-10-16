package main

type cliCmd struct {
	name        string
	description string
	callback    func(pd *pokedex) error
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
		"map": {
			name:        "map",
			description: "Get a map",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous map",
			callback:    cmdMapb,
		},
	}
}
