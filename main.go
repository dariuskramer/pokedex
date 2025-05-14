package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	Next     string
	Previous string
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

const (
	PokeApiLocationAreasURL = "https://pokeapi.co/api/v2/location-area/"
)

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func main() {
	const Prompt = "Pokedex > "
	config := &config{Next: PokeApiLocationAreasURL}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Prompt)
		if !scanner.Scan() {
			log.Fatalf("Scanner error: %v", scanner.Err())
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := supportedCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	var result []string

	for _, word := range strings.Fields(text) {
		result = append(result, strings.ToLower(word))
	}
	return result
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	usage := `
Welcome to the Pokedex!
Usage:

`
	for _, command := range supportedCommands {
		usage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	fmt.Println(usage)
	return nil
}

func internalCommandMap(url string, config *config) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer response.Body.Close()

	var payload LocationAreas
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&payload)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	for _, location := range payload.Results {
		fmt.Println(location.Name)
	}

	config.Next = payload.Next
	config.Previous = payload.Previous

	return nil
}

func commandMap(config *config) error {
	url := config.Next

	if url == "" {
		return fmt.Errorf("map: no next location areas")
	}

	err := internalCommandMap(url, config)
	if err != nil {
		return fmt.Errorf("map: %v", err)
	}

	return nil
}

func commandMapb(config *config) error {
	url := config.Previous

	if url == "" {
		return fmt.Errorf("mapb: no previous location areas")
	}

	err := internalCommandMap(url, config)
	if err != nil {
		return fmt.Errorf("mapb: %v", err)
	}

	return nil
}
