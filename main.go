package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dariuskramer/pokedex/internal/commands"
	"github.com/dariuskramer/pokedex/internal/pokeapi"
	"github.com/dariuskramer/pokedex/internal/repl"
)

func main() {
	const Prompt = "Pokedex > "
	config := &commands.CommandConfig{Next: pokeapi.PokeApiLocationAreasURL}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Prompt)
		if !scanner.Scan() {
			log.Fatalf("Scanner error: %v", scanner.Err())
		}
		input := repl.CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands.SupportedCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}
