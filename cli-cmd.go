package main

import (
	"maps"
	"slices"
)

type cliCmd struct {
	description string
	callback    func(cfg *config, args ...string) error
}

func cliCmds() map[string]cliCmd {
	return map[string]cliCmd{
		"cache": {
			description: "Show the cache",
			callback:    cmdCache,
		},
		"catch": {
			description: "Try to catch a pokemon",
			callback:    cmdCatch,
		},
		"exit": {
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"explore": {
			description: "Explore a location",
			callback:    cmdExplore,
		},
		"help": {
			description: "Display a help message",
			callback:    cmdHelp,
		},
		"inspect": {
			description: "Inspect your pokemon",
			callback:    cmdInspect,
		},
		"map": {
			description: "Get a map",
			callback:    cmdMap,
		},
		"mapb": {
			description: "Get previous map",
			callback:    cmdMapb,
		},
		"pokedex": {
			description: "Show your Pokedex",
			callback:    cmdPokedex,
		},
	}
}

func applyToSortedCmds(fn func(key string, cmd cliCmd)) {
	cmds := cliCmds()
	keys := []string{}
	for k := range maps.Keys(cmds) {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, key := range keys {
		fn(key, cmds[key])
	}
}
