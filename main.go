package main

import "fmt"

func main() {
	for _, word := range cleanInput(" Hello  World ") {
		fmt.Println(word)
	}
}
