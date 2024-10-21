package main

import (
	"io"
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestCmdMapb(t *testing.T) {
	cfg := newTestConfig(pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := `canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior

Previous
Next https://pokeapi.co/api/v2/location-area?offset=20&limit=20
`
	cfg.Previous = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cmdMapb(cfg)
	bytes, _ := io.ReadAll(cfg.output)
	got := string(bytes)
	if got != want {
		t.Fatalf("Want:\n%s\n, Got:\n%s", want, got)
	}
}

func TestCmdMapbPage2(t *testing.T) {
	cfg := newTestConfig(pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := `mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c

Previous https://pokeapi.co/api/v2/location-area?offset=0&limit=20
Next https://pokeapi.co/api/v2/location-area?offset=40&limit=20
`
	cfg.Previous = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
	cmdMapb(cfg)
	bytes, _ := io.ReadAll(cfg.output)
	got := string(bytes)
	if got != want {
		t.Fatalf("Want:\n%s\n, Got:\n%s", want, got)
	}
}

func TestCmdMapbInvalid(t *testing.T) {
	cfg := newTestConfig(pokeapi.NewPokeApi(pokecache.NewTestCache()))
	cfg.Previous = "lwkejffkejfwlkjef"
	err := cmdMapb(cfg)
	if err == nil {
		t.Fatalf("expected an error with invalid Next")
	}
}

func TestCmdMapbInvalidPage1(t *testing.T) {
	cfg := newTestConfig(pokeapi.NewPokeApi(pokecache.NewTestCache()))
	cfg.Previous = ""
	err := cmdMapb(cfg)
	if err == nil {
		t.Fatalf("expected an error when there is no Previous page")
	}
}

func TestCmdMapbNoNext(t *testing.T) {
	cfg := newTestConfig(pokeapi.NewPokeApi(pokecache.NewTestCache()))
	want := `team-rockets-castle-area
ultra-megalopolis-area
lush-jungle-all-areas
alola-route-17-all-areas

Previous https://pokeapi.co/api/v2/location-area?offset=1030&limit=20
Next
`
	cfg.Previous = "https://pokeapi.co/api/v2/location-area?offset=1050&limit=20"
	err := cmdMapb(cfg)
	if err != nil {
		t.Fatalf("Error with cmdMapb: %s", err.Error())
	}
	bytes, _ := io.ReadAll(cfg.output)
	got := string(bytes)
	if got != want {
		t.Fatalf("Want:\n%s\n, Got:\n%s", want, got)
	}
}
