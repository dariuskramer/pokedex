package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	var locations strings.Builder
	for _, location := range payload.Results {
		fmt.Println(location.Name)
		locations.WriteString(location.Name)
		locations.WriteByte('\n')
	}

	config.Next = payload.Next
	config.Previous = payload.Previous
	// if payload.Previous == "" {
	// 	config.Previous = url
	// } else {
	// 	config.Previous = payload.Previous
	// }
	removeLastNewline := len(locations.String()) - 1
	config.Cache.Add(url, []byte(locations.String()[:removeLastNewline]))

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
