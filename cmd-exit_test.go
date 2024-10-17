package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func TestCmdExit(t *testing.T) {
	var w bytes.Buffer
	want := "Exiting Pokedex\n"
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())

	cmdExit(config)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
