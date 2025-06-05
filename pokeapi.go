package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getData(url string, config *config, ptr any) (int, error) {
	cachedData, ok := config.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, ptr)
		if err != nil {
			return 0, fmt.Errorf("error: unable to unmarshal cached data")
		}
		return 0, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error: no response from server")
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return res.StatusCode, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(ptr)
	if err != nil {
		return 0, fmt.Errorf("error decoding response body")
	}

	marshaledData, err := json.Marshal(ptr)
	if err != nil {
		return 0, fmt.Errorf("error: unable to marshal data")
	}
	config.cache.Add(url, marshaledData)

	return res.StatusCode, nil
}

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

	_, err := getData(url, config, &config.pokeMap)
	if err != nil {
		return err
	}

	return nil
}

func getAreaData(config *config, area string) error {
	const baseURL = "https://pokeapi.co/api/v2/location-area/"
	url := baseURL + area

	resCode, err := getData(url, config, &config.pokeAreaData)
	if resCode == 404 {
		return fmt.Errorf("Location not found!")
	}
	if err != nil {
		return err
	}
	return nil
}

func getPokemonData(config *config, pokemon string) error {
	const baseURL = "https://pokeapi.co/api/v2/pokemon/"
	url := baseURL + pokemon

	resCode, err := getData(url, config, &config.pokemonData)
	if resCode == 404 {
		return fmt.Errorf("pokemon not found!")
	}
	if err != nil {
		return err
	}
	return nil
}
