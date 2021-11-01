package commands

import (
	"gitlab.cloudf.de/andre/trello-cli/config"
	"gitlab.cloudf.de/andre/trello-cli/trello"
)

type shoppingListCommand struct {
}

func init() {
	RegisterCommand(shoppingListCommand{})
}

func (command shoppingListCommand) GetInformation() (string, string) {
	return "version", "Print current version."
}

func (command shoppingListCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "shopping-list" || commandParams[0] == "sl" {
		return true
	}

	return false
}

func (command shoppingListCommand) Execute(commandParams []string) error {
	config, err := config.LoadConfig(config.GetConfigPath())
	if err != nil {
		return err
	}
	err, trelloService := trello.CreateNewTrelloService(
		config,
	)
	if err != nil {
		return err
	}

	err = trelloService.GetShoppingList()
	if err != nil {
		return err
	}

	return nil
}
