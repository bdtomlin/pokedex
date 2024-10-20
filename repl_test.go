package main

import (
	"bytes"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestStartRepl(t *testing.T) {
	stdin := strings.NewReader("exit\n")
	var w bytes.Buffer
	cfg := newConfig(stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))

	startRepl(cfg)
	want := "pokedex > \nExiting Pokedex os.Exit(0)"
	got := w.String()
	if strings.Contains(want, got) {
		t.Fatalf("Expected: %+v, Got: %+v", want, got)
	}
}

func TestNormalizeCmd(t *testing.T) {
	cases := []struct {
		input string
		cmd   string
		args  []string
	}{
		{input: "  val  ", cmd: "val", args: []string{}},
		{input: "  a b c  ", cmd: "a", args: []string{"b", "c"}},
	}
	for _, c := range cases {
		cmd, args := normalizeCmd(c.input)
		if cmd != c.cmd {
			t.Fatalf("Want: %s, Got: %s", c.cmd, cmd)
		}
		if slices.Compare(args, c.args) != 0 {
			t.Fatalf("Want args: %v, Got args: %v", c.args, args)
		}
	}
}

func TestExecCmd(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := "Error with command 'pokedex': Your Pokedex is empty!"
	execCommand("pokedex", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdInvalid(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := "invalid command\n"
	execCommand("invalidcmd", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdBlank(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := "\ninvalid command\n"
	execCommand("", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestPrintPrompt(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi(pokecache.NewTestCache()))
	printPrompt(cfg)

	want := "\npokedex > "
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
