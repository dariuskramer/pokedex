package commands

import (
	"errors"
	"fmt"
)

func CommandInspect(config *CommandConfig, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	pokemonToInspect := args[0]
	pokemonStat, ok := config.Pokedex[pokemonToInspect]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Print(pokemonStat.Formatter())

	return nil
}
