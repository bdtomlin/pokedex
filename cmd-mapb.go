package main

import (
	"errors"
	"fmt"

	"github.com/bdtomlin/pokedexcli/internal/pokeapi"
)

func cmdMapb(pd *pokedex) error {
	url := pd.Url
	if pd.Previous == "" {
		return errors.New("There is no previous map")
	}

	pMap, err := pokeapi.GetMap(url)
	if err != nil {
		return err
	}

	pd.Next = pMap.Next
	pd.Previous = pMap.Previous
	fmt.Fprintf(pd.output, "%+v", pMap)
	for _, result := range pMap.Results {
		fmt.Fprintln(pd.output, result.Name)
	}
	fmt.Fprintln(pd.output, "Previous", pd.Previous)
	fmt.Fprintln(pd.output, "Next", pd.Next)
	return nil
}
