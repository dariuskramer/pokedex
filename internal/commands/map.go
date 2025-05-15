package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dariuskramer/pokedex/internal/pokeapi"
)

func requestMap(url string, config *CommandConfig) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer response.Body.Close()

	var payload pokeapi.LocationAreas
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

func CommandMap(config *CommandConfig) error {
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

func CommandMapb(config *CommandConfig) error {
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
