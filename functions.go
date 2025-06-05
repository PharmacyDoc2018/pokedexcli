package main

import (
	"fmt"
	"math/rand"
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
			description: "displays a help message",
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
		"catch": {
			name:        "catch",
			description: "attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "lookup a pokemon's stats on the Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists the pokemon in the Pokedex",
			callback:    commandPokedex,
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
	config := &config{
		pokedex: map[string]pokemonData{},
	}
	return config
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

func commandCatch(c *config) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", c.lastInput[1])
	err := getPokemonData(c, c.lastInput[1])
	if err != nil {
		return err
	}
	baseXP := c.pokemonData.BaseExperience
	winningNumber := max(int(baseXP/30), 1)
	winningNumber = min(winningNumber, 5)
	randomNumber := rand.Intn(winningNumber) + 1
	if randomNumber == winningNumber {
		fmt.Printf("%s was caught!\n", c.lastInput[1])
		isNewEntry := enterDataInPokedex(c)
		if isNewEntry {
			fmt.Println("New entry entered into Pokedex!")
		}
	} else {
		fmt.Printf("%s escaped!\n", c.lastInput[1])
	}

	return nil
}

func commandInspect(c *config) error {
	pokemon := c.lastInput[1]
	entryExists := isInPokedex(c, pokemon)
	if !entryExists {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", c.pokedex[pokemon].Name)
		fmt.Printf("Height: %d\n", c.pokedex[pokemon].Height)
		fmt.Printf("Weight: %d\n", c.pokedex[pokemon].Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %d\n", c.pokedex[pokemon].Stats[0].BaseStat)
		fmt.Printf("  -attack: %d\n", c.pokedex[pokemon].Stats[1].BaseStat)
		fmt.Printf("  -defense: %d\n", c.pokedex[pokemon].Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %d\n", c.pokedex[pokemon].Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %d\n", c.pokedex[pokemon].Stats[4].BaseStat)
		fmt.Printf("  -speed: %d\n", c.pokedex[pokemon].Stats[5].BaseStat)
		fmt.Println("Types:")
		for i := 0; i < len(c.pokedex[pokemon].Types); i++ {
			fmt.Println("  -", c.pokedex[pokemon].Types[i].Type.Name)
		}
	}
	return nil
}

func commandPokedex(c *config) error {
	if len(c.pokedex) == 0 {
		fmt.Println("You have no pokemon registered on the Pokedex")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key := range c.pokedex {
		fmt.Println(" -", key)
	}
	return nil
}

func enterDataInPokedex(c *config) bool {
	pokemon := c.pokemonData.Name
	_, ok := c.pokedex[pokemon]
	if !ok {
		c.pokedex[pokemon] = c.pokemonData
		return true
	}
	return false
}

func isInPokedex(c *config, pokemon string) bool {
	_, ok := c.pokedex[pokemon]
	return ok
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

	switch command.name {
	case "map":
		fallthrough
	case "mapb":
		fallthrough
	case "exit":
		fallthrough
	case "explore":
		fallthrough
	case "catch":
		fallthrough
	case "inspect":
		fallthrough
	case "pokedex":
		err := command.callback(config)
		if err != nil {
			return err
		}
	default:
		command.callback(nil)
	}
	return nil
}
