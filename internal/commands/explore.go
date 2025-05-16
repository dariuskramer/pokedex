package commands

import (
	"errors"
	"fmt"
)

func CommandExplore(config *CommandConfig, args ...string) error {
	if len(args) == 0 {
		return errors.New("usage: explore <location...>")
	}

	location, err := config.PokeapiClient.GetLocationEncounters(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)

	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found!")
		return nil
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
