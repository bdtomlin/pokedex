package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdExplore(t *testing.T) {
	var w bytes.Buffer
	want := `- tentacool
- tentacruel
- staryu
- magikarp
- gyarados
- wingull
- pelipper
- shellos
- gastrodon
- finneon
- lumineon
`
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	locationArea := "canalave-city-area"

	cmdExplore(config, locationArea)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestCmdExploreNoLocation(t *testing.T) {
	var w bytes.Buffer
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))

	err := cmdExplore(config)
	if err == nil {
		t.Fatalf("Expected an error when there is no name")
	}

	err = cmdExplore(config, "")
	if err == nil {
		t.Fatalf("Expected an error when there is no name")
	}
}
