package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PharmacyDoc2018/pokedexcli/internal/pokecache"
)

func getCommands() commandMapList {
	commands := commandMapList{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "lists pokemon found in the area",
			callback:    commandExplore,
		},
	}
	return commands
}

func initREPL() *config {
	config := initConfig()
	config.cache, config.stop = initCache()
	return config
}

func initConfig() *config {
	var config config
	return &config
}

func initCache() (*pokecache.Cache, chan struct{}) {
	stop := make(chan struct{})
	pokeCache := pokecache.NewCache(10*time.Second, stop)
	return pokeCache, stop
}

func cleanInput(text string) []string {
	var textWords []string
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	firstPass := strings.Split(text, " ")

	for _, word := range firstPass {
		if word != "" {
			textWords = append(textWords, word)
		}
	}
	return textWords
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	close(c.stop)
	time.Sleep(500 * time.Millisecond)
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config) error {
	err := getLocationAreas(c, false)
	if err != nil {
		return err
	}
	for i := 0; i < len(c.pokeMap.Results); i++ {
		fmt.Println(c.pokeMap.Results[i].Name)
	}
	return nil
}

func commandMapB(c *config) error {
	err := getLocationAreas(c, true)
	if err != nil {
		return err
	}
	for i := 0; i < len(c.pokeMap.Results); i++ {
		fmt.Println(c.pokeMap.Results[i].Name)
	}
	return nil
}

func commandExplore(c *config) error {
	fmt.Printf("Exploring %s...\n", c.lastInput[1])
	err := getAreaData(c, c.lastInput[1])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for i := 0; i < len(c.pokeAreaData.PokemonEncounters); i++ {
		fmt.Printf(" - %s\n", c.pokeAreaData.PokemonEncounters[i].Pokemon.Name)
	}
	return nil
}

func cleanInputAndStore(c *config, input string) {
	c.lastInput = cleanInput(input)
}

func commandLookup(input string, commands commandMapList) (command cliCommand, err error) {
	for _, c := range commands {
		if input == c.name {
			return c, nil
		}
	}
	return cliCommand{}, fmt.Errorf("unknown command")
}

func commandLookupAndExecute(input string, commands commandMapList, config *config) error {
	cleanInputAndStore(config, input)
	command, err := commandLookup(config.lastInput[0], commands)
	if err != nil {
		return err
	}

	if command.name == "map" {
		err := command.callback(config)
		if err != nil {
			return err
		}
	} else if command.name == "mapb" {
		err := command.callback(config)
		if err != nil {
			return err
		}
	} else if command.name == "exit" {
		command.callback(config)
	} else if command.name == "explore" {
		command.callback(config)
	} else {
		command.callback(nil)
		return nil
	}
	return nil
}
