package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL     *string
	previousURL *string
}

type commandMap map[string]cliCommand
