package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	// Taken from https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_effort_value_yield_in_Generation_IX
	// Above this value I consider that the Pokémon is impossible to catch
	baseExperienceLimit = 300
)

func CommandCatch(config *CommandConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("usage: catch <pokémon>")
	}
	if len(args) > 1 {
		return errors.New("you can only catch one Pokémon at a time")
	}

	pokemon, err := config.PokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chanceToCatch := rand.Intn(baseExperienceLimit)
	if chanceToCatch >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	// Add to the Pokedex
	config.Pokedex[pokemon.Name] = pokemon

	return nil
}
