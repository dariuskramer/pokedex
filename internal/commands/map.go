package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

func fetchMap(url string, config *CommandConfig) error {
	// Is it in the cache?
	val, nextUrl, previousUrl, keyFound := config.Cache.Get(url)
	if keyFound {
		log.Println("cache hit!")
		fmt.Print(string(val))
		config.Next = nextUrl
		config.Previous = previousUrl
		return nil
	}

	// Not in the cache so fetch the data from the API
	var result pokeapi.LocationAreas
	err := pokeapi.Fetch(url, &result)
	if err != nil {
		return err
	}

	// Build the cached result
	var cachedResult strings.Builder
	for _, location := range result.Results {
		cachedResult.WriteString(location.Name)
		cachedResult.WriteByte('\n')
	}

	// Print the result
	fmt.Print(cachedResult.String())

	// Cache the result from the API
	config.Cache.Add(url, []byte(cachedResult.String()), result.Next, result.Previous)

	// Assign next/previous pagination
	config.Next = result.Next
	config.Previous = result.Previous

	return nil
}

func CommandMap(config *CommandConfig, args []string) error {
	var url string

	if config.Next != "" {
		url = config.Next
	} else {
		url = pokeapi.LocationAreasURL
	}

	err := fetchMap(url, config)
	if err != nil {
		return fmt.Errorf("map: %v", err)
	}

	return nil
}

func CommandMapb(config *CommandConfig, args []string) error {
	var url string

	if config.Previous != "" {
		url = config.Previous
	} else {
		url = pokeapi.LocationAreasURL
	}

	err := fetchMap(url, config)
	if err != nil {
		return fmt.Errorf("mapb: %v", err)
	}

	return nil
}
