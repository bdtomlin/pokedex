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
	return unmarshal(data, pMap)
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

	return unmarshal(data, la)
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

	return unmarshal(data, pok)
}

func unmarshal[T any](data []byte, strct T) (T, error) {
	if err := json.Unmarshal(data, &strct); err != nil {
		fmt.Println(string(data))
		return strct, fmt.Errorf("Error parsing json %w", err)
	}
	return strct, nil
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
