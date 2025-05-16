package commands

import (
	"fmt"
	"math/rand"

	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

const (
	// Taken from https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_effort_value_yield_in_Generation_IX
	// Above this value I consider that the Pokémon is impossible to catch
	baseExperienceLimit = 300
)

func CommandCatch(config *CommandConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: catch <pokémon>")
	}
	if len(args) > 1 {
		return fmt.Errorf("you can only catch one Pokémon at a time")
	}

	pokemonToCatch := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonToCatch)

	url := pokeapi.PokemonURL + "/" + pokemonToCatch

	// Fetch the data from the API
	var result pokeapi.Pokemon
	err := pokeapi.Fetch(url, &result)
	if err != nil {
		return fmt.Errorf("catch: %v", err)
	}

	chanceToCatch := rand.Intn(baseExperienceLimit)
	if chanceToCatch >= result.BaseExperience {
		fmt.Printf("%s was caught!\n", result.Name)

		// Add to the Pokedex
		config.Pokedex[result.Name] = result
	} else {
		fmt.Printf("%s escaped!\n", result.Name)
	}

	return nil
}
