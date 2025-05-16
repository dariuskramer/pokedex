package main

import (
	"time"

	"github.com/dariuskramer/pokedex/internal/commands"
	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

func main() {
	config := commands.CommandConfig{
		PokeapiClient: pokeapi.NewClient(5 * time.Second),
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}

	replLoop(&config)
}
