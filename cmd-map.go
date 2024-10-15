package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokeMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func cmdMap(w io.Writer) error {
	var pm pokeMap

	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		fmt.Fprintf(w, "Network error: %v", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pm); err != nil {
		fmt.Fprintf(w, "Json decode error: %v", err)
	}
	fmt.Fprintf(w, "%+v", pm)
	for _, result := range pm.Results {
		fmt.Fprintln(w, result.Name)
	}
	return nil
}
