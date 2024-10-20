package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

// This has a workaround to test randomness.
// I would like to look into other options for this.
func TestCmdCatchSuccess(t *testing.T) {
	var w bytes.Buffer
	want := "Throwing a Pokeball at pidgey...\npidgey was caught!\n"

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	cfg.RandIntFunc = func(i int) int { return 1 }
	cmdCatch(cfg, "pidgey")
	got := w.String()
	if got != want {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}

func TestCmdCatchFailure(t *testing.T) {
	var w bytes.Buffer
	want := "Throwing a Pokeball at pidgey...\npidgey escaped!\n"

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	cfg.RandIntFunc = func(i int) int { return 0 }
	cmdCatch(cfg, "pidgey")
	got := w.String()
	if got != want {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}

func TestCmdCatchNoName(t *testing.T) {
	var w bytes.Buffer

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))

	err := cmdCatch(cfg, "")
	if err == nil {
		t.Fatalf("Want an error when the name is blank")
	}

	err = cmdCatch(cfg)
	if err == nil {
		t.Fatalf("Want an error when the name is blank")
	}
}

func TestCmdCatchNonExistant(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	err := cmdCatch(cfg, "slkejfwerfwefjk")
	wantError := "Not Found"
	if err == nil {
		t.Errorf("Didn't get any errors wanted: %s", wantError)
		return
	}
	if !strings.Contains(err.Error(), wantError) {
		t.Errorf("Wanted error to contain: %s, Got error: %s", wantError, err.Error())
	}
}
