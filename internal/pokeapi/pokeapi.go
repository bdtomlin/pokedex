package pokeapi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
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

type NullCache struct{}

func (nc NullCache) Add(s string, b []byte)      {}
func (nc NullCache) Get(s string) ([]byte, bool) { return []byte{}, false }
func (nc NullCache) Dump() string                { return "nullCache" }

func NewPokeApi(maybeCache ...cache) *PokeApi {
	var cache cache
	if len(maybeCache) == 1 {
		cache = maybeCache[0]
	} else {
		cache = NullCache{}
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

	data, err := pApi.fetchData(url)
	if err != nil {
		return pMap, err
	}

	if err := json.Unmarshal(data, &pMap); err != nil {
		fmt.Println(string(data))
		return pMap, fmt.Errorf("Error parsing json1 %w", err)
	}
	return pMap, nil
}

func (pApi *PokeApi) GetLocation(area string) (LocationArea, error) {
	defer logResponseTime(time.Now())
	url := fullUrl("location-area", area)
	var la LocationArea

	data, err := pApi.fetchData(url)
	if err != nil {
		return la, err
	}

	if err := json.Unmarshal(data, &la); err != nil {
		fmt.Println(string(data))
		return la, fmt.Errorf("Error parsing json2 %w", err)
	}
	return la, nil
}

func (pApi *PokeApi) GetPokemon(name string) (Pokemon, error) {
	defer logResponseTime(time.Now())
	url := fullUrl("pokemon", name)
	var pok Pokemon

	data, err := pApi.fetchData(url)
	if err != nil {
		return pok, err
	}

	if err := json.Unmarshal(data, &pok); err != nil {
		fmt.Println("data", string(data))
		return pok, fmt.Errorf("Error parsing json3 %w", err)
	}
	return pok, nil
}

func (pApi *PokeApi) fetchData(url string) ([]byte, error) {
	data, ok := pApi.readDataFromCache(url)
	if !ok {
		apiData, err := pApi.readDataFromApi(url)
		if err != nil {
			return apiData, err
		}
		data = apiData
	}
	return data, nil
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
	var data []byte
	rawHttpRes, ok := pApi.Cache.Get(url)
	if !ok {
		return nil, false
	}
	r := bufio.NewReader(bytes.NewReader(rawHttpRes))
	res, err := http.ReadResponse(r, &http.Request{})
	if err != nil {
		return data, false
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, false
	}
	return data, true
}

func (pApi *PokeApi) readDataFromApi(url string) ([](byte), error) {
	var data []byte

	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("Request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode <= 599 && res.StatusCode >= 500 {
		return data, fmt.Errorf("Http error: %s", res.Status)
	}

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return data, fmt.Errorf("Error reading response: %w", err)
	}
	pApi.Cache.Add(url, dump)

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return data, fmt.Errorf("Http error: %s", res.Status)
	}
	return data, nil
}

func dataFromRawHttpRes(rawRes []byte) ([]byte, error) {
	var data []byte
	r := bufio.NewReader(bytes.NewReader(rawRes))
	res, err := http.ReadResponse(r, &http.Request{})
	if err != nil {
		return data, errors.New("error converting bytes to http response")
	}
	defer res.Body.Close()
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, errors.New("error reading response")
	}
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		return data, errors.New(res.Status)
	}
	return data, nil
}
