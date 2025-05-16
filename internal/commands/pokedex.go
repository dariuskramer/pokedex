package commands

import "fmt"

func CommandPokedex(config *CommandConfig, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
