package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocationAreas(config *config, isPrevious bool) error {
	var url string
	switch isPrevious {
	case false:
		if config.pokeMap.Next != nil {
			url = *config.pokeMap.Next
		} else {
			url = "https://pokeapi.co/api/v2/location-area"
		}
	case true:
		if config.pokeMap.Previous != nil {
			url = *config.pokeMap.Previous
		} else {
			return fmt.Errorf("you're on the first page")
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
	err = decoder.Decode(&config.pokeMap)
	if err != nil {
		return fmt.Errorf("error decoding response body")
	}

	return nil
}
