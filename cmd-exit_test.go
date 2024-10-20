package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdExit(t *testing.T) {
	var w bytes.Buffer
	want := "Exiting Pokedex\nos.Exit(0)\n"
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))

	cmdExit(config)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
	return
}
