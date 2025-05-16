package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

func CommandExplore(config *CommandConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: explore <location...>")
	}

	var result pokeapi.LocationsAreasEncounters
	for _, location := range args {
		url := pokeapi.LocationAreasURL + "/" + location

		fmt.Printf("Exploring %s...\n", location)

		// Is it in the cache?
		val, keyFound := config.Cache.Get(url)
		if keyFound {
			log.Println("cache hit!")
			fmt.Print(string(val))
			continue
		}

		// Not in the cache so fetch the data from the API
		err := pokeapi.Fetch(url, &result)
		if err != nil {
			return fmt.Errorf("explore: %v", err)
		}

		if len(result.PokemonEncounters) == 0 {
			fmt.Println("No Pokémon found!")
			continue
		}

		// Build the cached result
		var cachedResult strings.Builder
		cachedResult.WriteString("Found Pokémon:\n")
		for _, pokemon := range result.PokemonEncounters {
			s := fmt.Sprintf(" - %s\n", pokemon.Pokemon.Name)
			cachedResult.WriteString(s)
		}

		// Print the result
		fmt.Print(cachedResult.String())

		// Cache the result from the API
		config.Cache.Add(url, []byte(cachedResult.String()))
	}

	return nil
}
