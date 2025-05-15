package pokeapi

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
