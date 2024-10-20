package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdCache(t *testing.T) {
	var w bytes.Buffer
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	cmdCache(config)

	want := "This is a cache for testing. Cache is saved in the internal/pokecache/testcache directory\n"
	got := w.String()

	if want != got {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}
