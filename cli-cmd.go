package main

type cliCmd struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func cliCommands() map[string]cliCmd {
	return map[string]cliCmd{
		"cache": {
			name:        "cache",
			description: "Show the cache",
			callback:    cmdCache,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    cmdCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    cmdExplore,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    cmdHelp,
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
