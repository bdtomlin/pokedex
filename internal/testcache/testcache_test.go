package testcache

import (
	"slices"
	"testing"
)

func TestTestcache(t *testing.T) {
	tc := NewCache()
	want := []byte("whatever")

	tc.Add("mykey", want)
	got, _ := tc.Get("mykey")
	if slices.Compare(want, got) != 0 {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}
