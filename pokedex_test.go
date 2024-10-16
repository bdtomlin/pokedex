package main

import (
	"bytes"
	"testing"
)

func TestNewPokedex(t *testing.T) {
	var r bytes.Buffer
	var w bytes.Buffer
	pd := newPokedex(&r, &w)
	if pd.input != &r {
		t.Fatalf("Expected reader passed in to be in struct")
	}
	if pd.output != &w {
		t.Fatalf("Expected writer passed in to be in struct")
	}
	if pd.Next != "" {
		t.Fatalf("Expected Next to be a blank string")
	}
	if pd.Previous != "" {
		t.Fatalf("Expected Previous to be a blank string")
	}
}
