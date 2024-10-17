package main

import (
	"io"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	Name           string
	BaseExperience int
}

type config struct {
	input    io.Reader
	output   io.Writer
	Next     string
	Previous string
	PokeApi  *pokeapi.PokeApi
	Caught   []Pokemon
}

func newConfig(input io.Reader, output io.Writer, pokeApi *pokeapi.PokeApi) *config {
	return &config{
		input:   input,
		output:  output,
		PokeApi: pokeApi,
		Caught:  []Pokemon{},
	}
}
