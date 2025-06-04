package main

import "github.com/PharmacyDoc2018/pokedexcli/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type commandMapList map[string]cliCommand

type config struct {
	pokeMap mapConfig
	cache   *pokecache.Cache
	stop    chan struct{}
}

type mapConfig struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
