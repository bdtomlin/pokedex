package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCmdExit(t *testing.T) {
	var w bytes.Buffer
	want := "Exiting Pokedex\n"
	pd := newPokedex(os.Stdin, &w)

	cmdExit(&pd)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
