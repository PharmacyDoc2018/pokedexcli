package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getData(url string, config *config, ptr any) error {
	cachedData, ok := config.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, ptr)
		if err != nil {
			return fmt.Errorf("error: unable to unmarshal cached data")
		}
		return nil
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
	err = decoder.Decode(ptr)
	if err != nil {
		return fmt.Errorf("error decoding response body")
	}

	marshaledData, err := json.Marshal(ptr)
	if err != nil {
		return fmt.Errorf("error: unable to marshal data")
	}
	config.cache.Add(url, marshaledData)

	return nil
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

	err := getData(url, config, &config.pokeMap)
	if err != nil {
		return err
	}
	return nil
}

func getAreaData(config *config, area string) error {
	const baseURL = "https://pokeapi.co/api/v2/location-area"
	url := baseURL + "/" + area

	err := getData(url, config, &config.pokeAreaData)
	if err != nil {
		return err
	}
	return nil
}
