package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string) (PokeMap, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeMap{}, fmt.Errorf("Network error: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var pMap PokeMap
	if err := decoder.Decode(&pMap); err != nil {
		return PokeMap{}, fmt.Errorf("Json decode error: %w", err)
	}
	return pMap, nil
}
