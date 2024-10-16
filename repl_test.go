package main

import (
	"bytes"
	"os"
	"testing"
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
	pd := newPokedex(os.Stdin, &w)
	want := "Exiting Pokedex\n"
	execCommand("exit", &pd)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestExecCmdInvalid(t *testing.T) {
	var w bytes.Buffer
	pd := newPokedex(os.Stdin, &w)
	want := "invalid command\n"
	execCommand("invalidcmd", &pd)
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}

func TestPrintPrompt(t *testing.T) {
	var w bytes.Buffer
	pd := newPokedex(os.Stdin, &w)
	printPrompt(&pd)

	want := "\npokedex > "
	got := w.String()

	if got != want {
		t.Fatalf("Expected: %s, Got: %s", want, got)
	}
}
