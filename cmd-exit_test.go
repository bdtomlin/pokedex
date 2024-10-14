package main

import (
	"bytes"
	"testing"
)

func TestCmdExit(t *testing.T) {
	var w bytes.Buffer
	want := "Exiting Pokedex\n"

	cmdExit(&w)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
