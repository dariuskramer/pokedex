package commands

import (
	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

type CommandConfig struct {
	PokeapiClient    pokeapi.Client
	nextLocationsURL string
	prevLocationsURL string
	Pokedex          map[string]pokeapi.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*CommandConfig, ...string) error
}

var SupportedCommands map[string]CliCommand

func init() {
	SupportedCommands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 location areas in the Pokemon world",
			Callback:    CommandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 location areas in the Pokemon world",
			Callback:    CommandMapBackward,
		},
		"explore": {
			Name:        "explore",
			Description: "List all the Pokémon from a location area",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch Pokémon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect an already caught Pokémon",
			Callback:    CommandInspect,
		},
	}
}
