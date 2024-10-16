package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

type cache interface {
	Add(string, []byte)
	Get(string) ([]byte, bool)
}

func GetMap(url string, c cache) (PokeMap, error) {
	var pMap PokeMap

	data, ok := readDataFromCache(url, c)
	if !ok {
		apiData, err := readDataFromApi(url)
		if err != nil {
			return pMap, err
		}
		data = apiData
	}

	c.Add(url, data)
	if err := json.Unmarshal(data, &pMap); err != nil {
		return pMap, fmt.Errorf("Error parsing json %w", err)
	}
	return pMap, nil
}

func readDataFromCache(url string, c cache) ([](byte), bool) {
	data, ok := c.Get(url)
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

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("Error reading response body: %w", err)
	}
	return data, nil
}
