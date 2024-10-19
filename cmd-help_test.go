package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func TestCmdHelp(t *testing.T) {
	var w bytes.Buffer
	want := `
Welcome to Pokedex

Usage:

`
	appendSortedCmds := func(key string, cmd cliCmd) {
		want += fmt.Sprintf("%s: %s\n", key, cmd.description)
	}
	applyToSortedCmds(appendSortedCmds)

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	cmdHelp(cfg)
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
