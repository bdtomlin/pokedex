package pokeapi

import (
	"testing"

	"github.com/bdtomlin/pokedexcli/internal/pokecache"
)

func TestGetMap(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	pMap1, err := pApi.GetMap("")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(pMap1.Results) != 20 {
		t.Fatalf("Expected Results to have 20 results")
	}
	if pMap1.Previous != "" {
		t.Fatalf("Expected Previous to be blank")
	}

	pMap2, err := pApi.GetMap(pMap1.Next)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(pMap2.Results) != 20 {
		t.Fatalf("Expected Results to have 20 results")
	}
	if pMap2.Previous == "" {
		t.Fatalf("Expected Previous not to be blank")
	}
}

func TestGetMapErr(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	_, err := pApi.GetMap("weefkjwekjrfoj")
	if err == nil {
		t.Fatalf("expecting an error with invalid input")
	}
}

func TestGetLocation(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	_, err := pApi.GetLocation("")
	if err == nil {
		t.Fatalf("Expected an error for a blank location")
	}

	loc, err := pApi.GetLocation("canalave-city-area")
	if err != nil {
		t.Error(err)
	}
	if len(loc.PokemonEncounters) == 0 {
		t.Error("Expected some pokemon encounters")
	}
}

func TestGetLocationErr(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	_, err := pApi.GetLocation("weefkjwekjrfoj")
	if err == nil {
		t.Fatalf("expecting an error with invalid input")
	}
}

func TestGetPokemon(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)
	name := "pikachu"

	_, err := pApi.GetLocation("")
	if err == nil {
		t.Fatalf("Expected an error for a blank name")
	}

	pok, err := pApi.GetPokemon(name)
	if err != nil {
		t.Error(err)
	}
	if pok.Name != name {
		t.Fatalf("Want pokemon with name %s, Got: %s", name, pok.Name)
	}
}

func TestGetPokemonErr(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	_, err := pApi.GetPokemon("weefkjwekjrfoj")
	if err == nil {
		t.Fatalf("expecting an error with invalid input")
	}
}

func TestGetPokemonBlankName(t *testing.T) {
	testCache := pokecache.NewTestCache()
	pApi := NewPokeApi(testCache)

	_, err := pApi.GetPokemon("")
	if err == nil {
		t.Fatalf("expecting an error with blank input")
	}
}

func TestUnmarshal(t *testing.T) {
	str := struct {
		Name string `json:name`
	}{}
	data := []byte("abc")

	_, err := unmarshal(data, str)
	if err == nil {
		t.Fatalf("Expecting an error with invalid json")
	}
}

func TestNormalizeUrlOrPath(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "abc",
			output: baseUrl + "/abc",
		},
		{
			input:  baseUrl + "abc",
			output: baseUrl + "/abc",
		},
		{
			input:  baseUrl + "abc/",
			output: baseUrl + "/abc",
		},
		{
			input:  baseUrl + "/abc/",
			output: baseUrl + "/abc",
		},
		{
			input:  baseUrl + "abc/def/xyx",
			output: baseUrl + "/abc/def/xyx",
		},
		{
			input:  baseUrl,
			output: baseUrl,
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			want := c.output
			got := normalizeUrlOrPath(c.input)
			if got != want {
				t.Fatalf("Want: %s, Got: %s", want, got)
			}
		})
	}
}
