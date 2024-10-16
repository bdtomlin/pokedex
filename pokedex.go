package main

import "io"

type pokedex struct {
	input    io.Reader
	output   io.Writer
	Url      string
	Next     string
	Previous string
}

func newPokedex(input io.Reader, output io.Writer) pokedex {
	return pokedex{
		input:  input,
		output: output,
		Url:    "https://pokeapi.co/api/v2/location-area",
	}
}
