package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestNormalizeCmd(t *testing.T) {
	tm := map[string]string{
		"  val  ":  "val",
		"a b c   ": "a",
	}

	for k, v := range tm {
		want := v
		got := normalizeCmd(k)
		if got != want {
			t.Fatalf("Want: %s, Got: %s", want, got)
		}
	}
}

func TestExecCmd(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokecache.NewCache())
	want := "Exiting Pokedex\n"
	execCommand("exit", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdInvalid(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokecache.NewCache())
	want := "invalid command\n"
	execCommand("invalidcmd", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdBlank(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokecache.NewCache())
	want := "\ninvalid command\n"
	execCommand("", cfg)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestPrintPrompt(t *testing.T) {
	var w bytes.Buffer
	cfg := newConfig(os.Stdin, &w, pokecache.NewCache())
	printPrompt(cfg)

	want := "\npokedex > "
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
