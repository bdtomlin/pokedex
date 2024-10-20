package main

import (
	"io"
	"math/rand"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

type Stats []struct {
	BaseStat int
	Effort   int
	Stat     struct {
		Name string
		URL  string
	}
}

type Types []struct {
	Slot int
	Type struct {
		Name string
		URL  string
	}
}

type config struct {
	input       io.Reader
	output      io.Writer
	Next        string
	Previous    string
	PokeApi     *pokeapi.PokeApi
	Caught      map[string]pokeapi.Pokemon
	RandIntFunc func(int) int
}

func newConfig(input io.Reader, output io.Writer, pokeApi *pokeapi.PokeApi) *config {
	return &config{
		input:       input,
		output:      output,
		PokeApi:     pokeApi,
		Caught:      map[string]pokeapi.Pokemon{},
		RandIntFunc: randIntFunc,
	}
}

func randIntFunc(i int) int {
	return rand.Intn(i)
}
