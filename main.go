package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()

	pokeScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	for pokeScanner.Scan() {
		input := pokeScanner.Text()
		err := commandLookupAndExecute(input, commands)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Pokedex > ")
	}

}
