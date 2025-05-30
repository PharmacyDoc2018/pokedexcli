package main

import (
	"fmt"
	"os"
	"strings"
)

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

func commandLookup(input string, commands commandMap) {
	commandFound := false
	for _, c := range commands {
		if strings.ToLower(input) == c.name {
			commandFound = true
			c.callback()
			break
		}
	}
	if !commandFound {
		fmt.Println("Unknown command")
	}
}
