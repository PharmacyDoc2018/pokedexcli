package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := commandMap{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	pokeScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	for pokeScanner.Scan() {
		input := pokeScanner.Text()
		commandLookup(input, commands)
		fmt.Printf("Pokedex > ")
	}

}
