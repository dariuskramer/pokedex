package commands

import "fmt"

func CommandHelp(config *CommandConfig) error {
	usage := `
Welcome to the Pokedex!
Usage:

`
	for _, command := range SupportedCommands {
		usage += fmt.Sprintf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Println(usage)
	return nil
}
