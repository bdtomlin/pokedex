package main

import (
	"fmt"
)

func cmdMap(cfg *config, args ...string) error {
	path := ""
	if cfg.Next != "" {
		path = cfg.Next
	}

	pMap, err := cfg.PokeApi.GetMap(path)
	if err != nil {
		return err
	}

	cfg.Next = pMap.Next
	cfg.Previous = pMap.Previous
	for _, result := range pMap.Results {
		fmt.Fprintln(cfg.output, result.Name)
	}

	fmt.Fprintln(cfg.output)

	if cfg.Previous == "" {
		fmt.Fprintln(cfg.output, "Previous")
	} else {
		fmt.Fprintln(cfg.output, "Previous", cfg.Previous)
	}

	if cfg.Next == "" {
		fmt.Fprintln(cfg.output, "Next")
	} else {
		fmt.Fprintln(cfg.output, "Next", cfg.Next)
	}
	return nil
}
