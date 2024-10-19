package main

import (
	"slices"
	"testing"
)

func TestApplyToSortedCmds(t *testing.T) {
	wantSorted := []string{}

	for k := range cliCmds() {
		wantSorted = append(wantSorted, k)
	}
	slices.Sort(wantSorted)

	gotSorted := []string{}

	orderKeys := func(key string, cmd cliCmd) {
		gotSorted = append(gotSorted, key)
	}
	applyToSortedCmds(orderKeys)

	if slices.Compare(wantSorted, gotSorted) != 0 {
		t.Fatalf("Expected keys to be sorted. Want: %v, Got: %v", wantSorted, gotSorted)
	}
}
