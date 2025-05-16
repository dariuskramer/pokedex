package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c Client) ListLocationAreas(pageURL string) (LocationAreas, error) {
	url := baseUrl + "/location-area"
	if pageURL != "" {
		url = pageURL
	}

	// Is it in the cache?
	cachedResponse, keyFound := c.cache.Get(url)
	if keyFound {
		log.Println("cache hit!")

		var locations LocationAreas
		err := json.Unmarshal(cachedResponse, &locations)
		if err != nil {
			return LocationAreas{}, err
		}
		return locations, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreas{}, err
	}
	defer response.Body.Close()

	var locations LocationAreas
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	// Cache the JSON response from the API
	c.cache.Add(url, []byte(data))

	return locations, nil
}
