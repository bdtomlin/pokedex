package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/testcache"
)

func TestCmdCache(t *testing.T) {
	var w bytes.Buffer
	config := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(testcache.NewCache()))
	cmdCache(config)

	want := "this is a cache for testing\n"
	got := w.String()

	if want != got {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}
