package main

import (
	"bytes"
	"os"
	"slices"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

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
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	want := "Exiting Pokedex\n"
	execCommand("exit", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdInvalid(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	want := "invalid command\n"
	execCommand("invalidcmd", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdBlank(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	want := "\ninvalid command\n"
	execCommand("", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestPrintPrompt(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokeapi.NewPokeApi())
	printPrompt(cfg)

	want := "\npokedex > "
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
