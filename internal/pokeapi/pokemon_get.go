package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	// Is it in the cache?
	cachedResponse, keyFound := c.cache.Get(url)
	if keyFound {
		log.Println("cache hit!")

		var pokemon Pokemon
		err := json.Unmarshal(cachedResponse, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	var pokemon Pokemon
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// Cache the JSON response from the API
	c.cache.Add(url, []byte(data))

	return pokemon, nil
}
