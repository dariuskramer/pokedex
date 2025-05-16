package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

func requestMap(url string, config *CommandConfig) error {
	val, keyFound := config.Cache.Get(url)
	if keyFound {
		log.Println("cache hit!")
		fmt.Println(string(val))
		return nil
	}

	var result pokeapi.LocationAreas
	err := pokeapi.Fetch(url, &result)
	if err != nil {
		return err
	}

	var locations strings.Builder
	for _, location := range result.Results {
		fmt.Println(location.Name)
		locations.WriteString(location.Name)
		locations.WriteByte('\n')
	}

	config.Next = result.Next
	config.Previous = result.Previous
	// if payload.Previous == "" {
	// 	config.Previous = url
	// } else {
	// 	config.Previous = payload.Previous
	// }
	removeLastNewline := len(locations.String()) - 1
	config.Cache.Add(url, []byte(locations.String()[:removeLastNewline]))

	return nil
}

func CommandMap(config *CommandConfig, args []string) error {
	url := config.Next

	if url == "" {
		return fmt.Errorf("map: no next location areas")
	}

	err := requestMap(url, config)
	if err != nil {
		return fmt.Errorf("map: %v", err)
	}

	return nil
}

func CommandMapb(config *CommandConfig, args []string) error {
	url := config.Previous

	if url == "" {
		return fmt.Errorf("mapb: no previous location areas")
	}

	err := requestMap(url, config)
	if err != nil {
		return fmt.Errorf("mapb: %v", err)
	}

	return nil
}
