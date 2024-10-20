package pokecache

import (
	"os"
	"slices"
	"testing"
)

func TestTestcache(t *testing.T) {
	cases := []string{
		"https://pokeapi.co/api/v2/location-area",
		"https://pokeapi.co/api/v2/location-area/lwlekjfwelkje",
	}
	for _, url := range cases {
		tc := NewTestCache()
		bytes, err := tc.Get(url)

		filePath := getCacheDir() + "/" + fileNameHash(url)
		onDisk, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Expected a cache file written to %s", filePath)
		}
		if slices.Compare(bytes, onDisk) == 0 {
			t.Fatalf("The cache file appears incorrect %s", filePath)
		}

	}
}
