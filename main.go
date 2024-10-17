package main

import (
	"os"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache()
	pokeApi := pokeapi.NewPokeApi(cache)
	config := newConfig(os.Stdin, os.Stdout, pokeApi)
	startRepl(config)
}
