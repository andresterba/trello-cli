package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andresterba/trello-cli/commands"
	"github.com/andresterba/trello-cli/config"
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

	config, err := config.LoadConfig(config.GetConfigPath())
	checkForError(err)

	params := args[1:]

	commandExecuted := false

	determineAndSetContext(params, config)
	err = config.WriteConfig()
	checkForError(err)

	params = args[3:]

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

func determineAndSetContext(params []string, config *config.Config) {
	if params[0] == "--context" {
		context := params[1]
		if checkIfContextIsValid(context) {
			config.DefaultContext = context
		}
	}
}

func checkIfContextIsValid(context string) bool {
	if context == commands.PersonalContext || context == commands.WorkContext {
		return true
	}

	return false
}
