package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const baseUrl = "https://pokeapi.co/api/v2"

type PokeApi struct {
	Cache cache
}

type cache interface {
	Get(string) ([]byte, error)
	Dump() string
}

func NewPokeApi(c cache) *PokeApi {
	return &PokeApi{
		Cache: c,
	}
}

type PokeMap struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (pApi *PokeApi) GetMap(url string) (PokeMap, error) {
	if url == "" {
		url = "/location-area"
	}
	url = normalizeUrlOrPath(url)
	var pMap PokeMap

	data, err := pApi.Cache.Get(url)
	if err != nil {
		return pMap, err
	}

	if err := json.Unmarshal(data, &pMap); err != nil {
		fmt.Println(string(data))
		return pMap, fmt.Errorf("Error parsing json %w", err)
	}
	return pMap, nil
}

func (pApi *PokeApi) GetLocation(area string) (LocationArea, error) {
	var la LocationArea

	if area == "" {
		return la, errors.New(`An area is required, got "" `)
	}
	url := normalizeUrlOrPath("location-area/" + area)

	data, err := pApi.Cache.Get(url)
	if err != nil {
		return la, err
	}

	if err := json.Unmarshal(data, &la); err != nil {
		fmt.Println(string(data))
		return la, fmt.Errorf("Error parsing json %w", err)
	}
	return la, nil
}

func (pApi *PokeApi) GetPokemon(name string) (Pokemon, error) {
	var pok Pokemon

	if name == "" {
		return pok, errors.New(`A name is required, got "" `)
	}
	url := normalizeUrlOrPath("pokemon/" + name)

	data, err := pApi.Cache.Get(url)
	if err != nil {
		return pok, err
	}

	if err := json.Unmarshal(data, &pok); err != nil {
		fmt.Println("data", string(data))
		return pok, fmt.Errorf("Error parsing json %w", err)
	}
	return pok, nil
}

func normalizeUrlOrPath(url string) string {
	if url == baseUrl {
		return url
	}

	url = strings.Replace(url, baseUrl, "", 1)
	parts := []string{baseUrl, url}

	for i, part := range parts {
		parts[i] = strings.Trim(part, "/")
	}
	return strings.Join(parts, "/")
}
