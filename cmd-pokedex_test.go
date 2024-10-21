package main

import (
	"io"
	"strings"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdPokedex(t *testing.T) {
	pApi := pokeapi.NewPokeApi(pokecache.NewTestCache())
	cfg := newTestConfig(pApi)
	p, _ := pApi.GetPokemon("pikachu")
	p2, _ := pApi.GetPokemon("pidgey")
	capture(cfg, p)
	capture(cfg, p2)
	cmdPokedex(cfg)
	bytes, err := io.ReadAll(cfg.output)
	if err != nil {
		t.Fatalf("Error reading output: %s", err.Error())
	}
	got := string(bytes)
	if !strings.Contains(got, "  - pikachu") {
		t.Fatal("Expected string to contain -pikachu")
	}
	if !strings.Contains(got, "  - pidgey") {
		t.Fatal("Expected string to contain -pikachu")
	}
}

func TestCmdPokedexEmpty(t *testing.T) {
	pApi := pokeapi.NewPokeApi(pokecache.NewTestCache())
	cfg := newTestConfig(pApi)
	err := cmdPokedex(cfg)
	if err == nil {
		t.Fatalf("expecting an error with empty pokedex")
	}
}

func TestCmdPokedexExtraArgs(t *testing.T) {
	pApi := pokeapi.NewPokeApi(pokecache.NewTestCache())
	cfg := newTestConfig(pApi)
	err := cmdPokedex(cfg, "whatever")
	if err == nil {
		t.Fatalf("expecting an error when there are extra args")
	}
}
