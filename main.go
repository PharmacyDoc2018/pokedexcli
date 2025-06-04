package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()
	config := initREPL()

	pokeScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	for pokeScanner.Scan() {
		input := pokeScanner.Text()
		err := commandLookupAndExecute(input, commands, config)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\n")
		fmt.Printf("Pokedex > ")
	}

}
