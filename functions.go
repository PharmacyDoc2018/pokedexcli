package main

import (
	"fmt"
	"os"
	"strings"
)

func getCommands() commandMap {
	commands := commandMap{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return commands
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandLookup(input string, commands commandMap) (command cliCommand, err error) {
	for _, c := range commands {
		if strings.ToLower(input) == c.name {
			return c, nil
		}
	}
	return cliCommand{}, fmt.Errorf("unknown command")
}

func commandLookupAndExecute(input string, commands commandMap) error {
	command, err := commandLookup(input, commands)
	if err != nil {
		return err
	}
	command.callback()
	return nil
}
