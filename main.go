package main

import (
	"os"
	"time"

	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func main() {
	config := newConfig(os.Stdin, os.Stdout, pokecache.NewCache(10*time.Second))
	startRepl(config)
}
