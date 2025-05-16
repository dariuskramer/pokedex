package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	LocationAreasURL = "https://pokeapi.co/api/v2/location-area/"
	PokemonURL       = "https://pokeapi.co/api/v2/pokemon/"
)

func Fetch(url string, result any) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationsAreasEncounters struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) Formatter() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "Name: %s\n", p.Name)
	fmt.Fprintf(&builder, "Base experience: %d\n", p.BaseExperience)

	builder.WriteString("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Fprintf(&builder, "  - %s\n", stat.Type.Name)
	}

	builder.WriteString("Types:\n")
	for _, types := range p.Types {
		fmt.Fprintf(&builder, "  - %s\n", types.Type.Name)
	}

	return builder.String()
}
