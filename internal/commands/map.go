package commands

import (
	"errors"
	"fmt"
)

func CommandMapForward(config *CommandConfig, args ...string) error {
	url := config.nextLocationsURL
	response, err := config.PokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}

	config.nextLocationsURL = response.Next
	config.prevLocationsURL = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandMapBackward(config *CommandConfig, args ...string) error {
	if config.prevLocationsURL == "" {
		return errors.New("already at the first page")
	}

	url := config.prevLocationsURL
	response, err := config.PokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}

	config.nextLocationsURL = response.Next
	config.prevLocationsURL = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
