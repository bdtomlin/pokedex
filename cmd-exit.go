package main

import (
	"fmt"
	"os"
)

func cmdExit(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.output, "Exiting Pokedex")

	// This is to make the func testable
	defer func() {
		if r := recover(); r == "unexpected call to os.Exit(0) during test" {
			fmt.Fprintln(cfg.output, "os.Exit(0)")
		}
	}()

	os.Exit(0)
	return nil
}
