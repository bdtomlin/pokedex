package main

import (
	"fmt"
)

func cmdCache(cfg *config) error {
	fmt.Fprintln(cfg.output, cfg.Cache.Dump())
	return nil
}
