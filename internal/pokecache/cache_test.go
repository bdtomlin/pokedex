package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheGet(t *testing.T) {
	cases := []struct {
		url     string
		isError bool
	}{
		{
			url:     "https://pokeapi.co/api/v2/location-area",
			isError: false,
		},
		{
			url:     "https://pokeapi.co/api/v2/location-area/thisisa404",
			isError: true,
		},
	}
	const interval = 5 * time.Second
	for _, c := range cases {
		t.Run(fmt.Sprintf("testing %v", c.url), func(t *testing.T) {
			cache := NewCache(interval)
			_, err := cache.Get(c.url)
			if c.isError && err == nil {
				t.Fatal("Expected and error")
			}
			if !c.isError && err != nil {
				t.Fatal("Expected no error")
			}
			_, ok := cache.GetRaw(c.url)
			if !ok {
				t.Fatalf("cache is missing %s", c.url)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	testUrl := "https://pokeapi.co/api/v2/location-area"
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Get(testUrl)

	_, ok := cache.GetRaw(testUrl)
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	time.Sleep(waitTime)

	_, ok = cache.GetRaw(testUrl)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	return
}
