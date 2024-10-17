package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type PokeApi struct {
	Cache cache
}

type cache interface {
	Add(string, []byte)
	Get(string) ([]byte, bool)
	Dump() string
}

type nullCache struct{}

func (nc nullCache) Add(s string, b []byte)      {}
func (nc nullCache) Get(s string) ([]byte, bool) { return []byte{}, false }
func (nc nullCache) Dump() string                { return "nullCache" }

func NewPokeApi(maybeCache ...cache) *PokeApi {
	var cache cache
	if len(maybeCache) == 1 {
		cache = maybeCache[0]
	} else {
		cache = nullCache{}
	}
	return &PokeApi{
		Cache: cache,
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

func logResponseTime(start time.Time) {
	end := time.Now()
	fmt.Printf("\nResponse time: %v\n", end.Sub(start))
}

func (pApi *PokeApi) GetMap(url string) (PokeMap, error) {
	defer logResponseTime(time.Now())
	if url == "" {
		url = "/location-area"
	}
	url = normalizeUrlOrPath(url)

	var pMap PokeMap

	data, ok := pApi.readDataFromCache(url)
	if !ok {
		apiData, err := readDataFromApi(url)
		if err != nil {
			return pMap, err
		}
		data = apiData
	}

	pApi.Cache.Add(url, data)
	if err := json.Unmarshal(data, &pMap); err != nil {
		return pMap, fmt.Errorf("Error parsing json %w", err)
	}
	return pMap, nil
}

func (pApi *PokeApi) GetLocation(area string) (LocationArea, error) {
	defer logResponseTime(time.Now())
	url := fullUrl("location-area", area)
	var la LocationArea

	data, ok := pApi.readDataFromCache(url)
	if !ok {
		apiData, err := readDataFromApi(url)
		if err != nil {
			return la, err
		}
		data = apiData
	}

	pApi.Cache.Add(url, data)
	if err := json.Unmarshal(data, &la); err != nil {
		return la, fmt.Errorf("Error parsing json %w", err)
	}
	return la, nil
}

func (pApi *PokeApi) GetPokemon(name string) (Pokemon, error) {
	defer logResponseTime(time.Now())
	url := fullUrl("pokemon", name)
	var pok Pokemon
	fmt.Println(url)

	data, ok := pApi.readDataFromCache(url)
	if !ok {
		apiData, err := readDataFromApi(url)
		if err != nil {
			return pok, err
		}
		data = apiData
	}

	pApi.Cache.Add(url, data)
	if err := json.Unmarshal(data, &pok); err != nil {
		return pok, fmt.Errorf("Error parsing json %w", err)
	}
	return pok, nil
}

func normalizeUrlOrPath(url string) string {
	url = strings.Replace(url, baseUrl, "", 1)
	return fullUrl(url)
}

func fullUrl(parts ...string) string {
	parts = append([]string{baseUrl}, parts...)
	for i, part := range parts {
		parts[i] = strings.Trim(part, "/")
	}
	return strings.Join(parts, "/")
}

func (pApi *PokeApi) readDataFromCache(url string) ([](byte), bool) {
	data, ok := pApi.Cache.Get(url)
	if !ok {
		return nil, false
	}
	return data, true
}

func readDataFromApi(url string) ([](byte), error) {
	var data []byte

	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("Network error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return data, fmt.Errorf("Http error: %s", res.Status)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("Error reading response body: %w", err)
	}
	return data, nil
}
