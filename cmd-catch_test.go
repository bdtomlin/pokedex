package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/testcache"
)

// This has a workaround to test randomness.
// I would like to look into other options for this.
func TestCmdCatchSuccess(t *testing.T) {
	var w bytes.Buffer
	want := "Throwing a Pokeball at pidgey...\npidgey was caught!\n"

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(testcache.NewCache()))
	cfg.testing = true
	cfg.testingRandSuccess = true
	cmdCatch(cfg, "pidgey")
	got := w.String()
	if got != want {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}

func TestCmdCatchFailure(t *testing.T) {
	var w bytes.Buffer
	want := "Throwing a Pokeball at pidgey...\npidgey escaped!\n"

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(testcache.NewCache()))
	cfg.testing = true
	cfg.testingRandSuccess = false
	cmdCatch(cfg, "pidgey")
	got := w.String()
	if got != want {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}

func TestCmdCatchNonExistentName(t *testing.T) {
	var w bytes.Buffer
	want := "Throwing a Pokeball at pidgey...\npidgey escaped!\n"

	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(testcache.NewCache()))
	cfg.testing = true
	cfg.testingRandSuccess = false
	cmdCatch(cfg, "pidgey")
	got := w.String()
	if got != want {
		t.Fatalf("Want: %s, Got: %s", want, got)
	}
}

func TestCmdCatchNoName(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
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
