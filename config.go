package main

import (
	"io"

	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

type config struct {
	input    io.Reader
	output   io.Writer
	Url      string
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

func newConfig(input io.Reader, output io.Writer, cache *pokecache.Cache) *config {
	return &config{
		input:  input,
		output: output,
		Url:    "https://pokeapi.co/api/v2/location-area",
		Cache:  cache,
	}
}
