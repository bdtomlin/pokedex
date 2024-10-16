package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdHelp(t *testing.T) {
	var w bytes.Buffer
	want := `
Welcome to Pokedex

Usage:

cache: Show the cache
exit: Exit the Pokedex
help: Displays a help message
map: Get a map
mapb: Get previous map
`
	cfg := newConfig(os.Stdin, &w, pokecache.NewCache())
	cmdHelp(cfg)
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
