package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c Client) GetLocationEncounters(locationName string) (LocationAreaEncounters, error) {
	url := baseUrl + "/location-area/" + locationName

	// Is it in the cache?
	cachedResponse, keyFound := c.cache.Get(url)
	if keyFound {
		log.Println("cache hit!")

		var encounters LocationAreaEncounters
		err := json.Unmarshal(cachedResponse, &encounters)
		if err != nil {
			return LocationAreaEncounters{}, err
		}
		return encounters, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaEncounters{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreaEncounters{}, err
	}
	defer response.Body.Close()

	var encounters LocationAreaEncounters
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaEncounters{}, err
	}

	err = json.Unmarshal(data, &encounters)
	if err != nil {
		return LocationAreaEncounters{}, err
	}

	// Cache the JSON response from the API
	c.cache.Add(url, []byte(data))

	return encounters, nil
}
