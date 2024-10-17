package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func TestCmdHelp(t *testing.T) {
	var w bytes.Buffer
	want := `
Welcome to Pokedex

Usage:

cache: Show the cache
catch: Try to catch a pokemon
exit: Exit the Pokedex
explore: Explore a location
help: Display a help message
map: Get a map
mapb: Get previous map
`
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	cmdHelp(cfg)
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
