package main

import (
	"bytes"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestNewConfig(t *testing.T) {
	var r bytes.Buffer
	var w bytes.Buffer
	cfg := newConfig(&r, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	if cfg.input != &r {
		t.Fatalf("Expected reader passed in to be in struct")
	}
	if cfg.output != &w {
		t.Fatalf("Expected writer passed in to be in struct")
	}
	if cfg.Next != "" {
		t.Fatalf("Expected Next to be a blank string")
	}
	if cfg.Previous != "" {
		t.Fatalf("Expected Previous to be a blank string")
	}
}
