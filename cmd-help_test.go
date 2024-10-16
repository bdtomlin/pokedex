package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCmdHelp(t *testing.T) {
	var w bytes.Buffer
	want := `
Welcome to Pokedex

Usage:

exit: Exit the Pokedex
help: Displays a help message
map: Get a map
mapb: Get previous map
`
	pd := newPokedex(os.Stdin, &w)
	cmdHelp(&pd)
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
