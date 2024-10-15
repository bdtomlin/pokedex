package main

import (
	"bytes"
	"testing"
)

func TestCmdHelp(t *testing.T) {
	var w bytes.Buffer
	want := `
Welcome to Pokedex

Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	cmdHelp(&w)
	got := w.String()
	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
