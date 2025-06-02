package main

import (
	"fmt"
	"os"
	"strings"
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
	}
	return commands
}

func initConfig() *config {
	var config config
	return &config
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

func commandExit(any) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(any) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c any) error {
	value, ok := c.(*config)
	if !ok {
		return fmt.Errorf("error: mapConfig pointer needed")
	}
	err := getLocationAreas(&value.pokeMap, false)
	if err != nil {
		return err
	}
	for i := 0; i < len(value.pokeMap.Results); i++ {
		fmt.Println(value.pokeMap.Results[i].Name)
	}
	return nil
}

func commandMapB(c any) error {
	value, ok := c.(*config)
	if !ok {
		return fmt.Errorf("error: mapConfig pointer needed")
	}
	err := getLocationAreas(&value.pokeMap, true)
	if err != nil {
		return err
	}
	for i := 0; i < len(value.pokeMap.Results); i++ {
		fmt.Println(value.pokeMap.Results[i].Name)
	}
	return nil
}

func commandLookup(input string, commands commandMapList) (command cliCommand, err error) {
	for _, c := range commands {
		if strings.ToLower(input) == c.name {
			return c, nil
		}
	}
	return cliCommand{}, fmt.Errorf("unknown command")
}

func commandLookupAndExecute(input string, commands commandMapList, config any) error {
	command, err := commandLookup(input, commands)
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
	} else {
		command.callback(nil)
		return nil
	}
	return nil
}
