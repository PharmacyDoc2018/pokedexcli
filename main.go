package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pokeScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	for pokeScanner.Scan() {
		input := pokeScanner.Text()
		response := cleanInput(input)[0]
		fmt.Printf("Your command was: %s\n", response)
		fmt.Printf("Pokedex > ")
	}

}
