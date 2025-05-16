package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dariuskramer/pokedex/internal/commands"
)

func replLoop(config *commands.CommandConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			log.Fatalf("Scanner error: %v", scanner.Err())
		}

		input := CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands.SupportedCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		var args []string
		if len(input) > 1 {
			args = input[1:]
		}

		err := command.Callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
