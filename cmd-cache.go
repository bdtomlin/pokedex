package main

import (
	"fmt"
)

func cmdCache(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.output, cfg.PokeApi.Cache.Dump())
	return nil
}
