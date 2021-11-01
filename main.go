package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.cloudf.de/andre/trello-cli/commands"
)

func checkForError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func showHelp() {
	fmt.Println(`
trello-cli [command] [options]
commands:`)

	for _, command := range commands.GetRegisteredCommands() {
		name, description := command.GetInformation()
		fmt.Printf("    %s - %s\n", name, description)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		showHelp()
		os.Exit(1)
	}

	params := args[1:]

	commandExecuted := false

	for _, command := range commands.GetRegisteredCommands() {
		if command.IsForCommand(params) {
			err := command.Execute(params)
			checkForError(err)
			commandExecuted = true

			break
		}
	}

	if !commandExecuted {
		showHelp()
		os.Exit(1)
	}
}
