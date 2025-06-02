package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocationAreas(config *locationAreaRes, isPrevious bool) error {
	var url string
	switch isPrevious {
	case false:
		if config.Next != nil {
			url = *config.Next
		} else {
			url = "https://pokeapi.co/api/v2/location-area"
		}
	case true:
		if config.Previous != nil {
			url = *config.Previous
		} else {
			return fmt.Errorf("error: no previous map")
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error: no response from server")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&config)
	if err != nil {
		return fmt.Errorf("error decoding response body")
	}

	return nil
}
