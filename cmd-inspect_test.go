package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdInspect(t *testing.T) {
	var w bytes.Buffer
	want := `Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  -electric
`

	c := pokecache.NewTestCache()
	r, _ := io.Pipe()
	cfg := newConfig(r, &w, pokeapi.NewPokeApi(c))
	pApi := pokeapi.NewPokeApi(c)
	p, _ := pApi.GetPokemon("pikachu")

	capture(cfg, p)
	cmdInspect(cfg, "pikachu")
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestCmdInspectWithNoName(t *testing.T) {
	var w bytes.Buffer
	c := pokecache.NewTestCache()
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(c))

	err := cmdInspect(cfg, "")
	if err == nil {
		t.Fatalf("Expected an error with blank name")
	}
	err = cmdInspect(cfg)
	if err == nil {
		t.Fatalf("Expected an error with no name")
	}
}

func TestCmdInspectWithUncaughtName(t *testing.T) {
	var w bytes.Buffer
	c := pokecache.NewTestCache()
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(c))

	err := cmdInspect(cfg, "purdydoo")
	if err == nil {
		t.Fatalf("Expected an error with uncaught name")
	}
}
