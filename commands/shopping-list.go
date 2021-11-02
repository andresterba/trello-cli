package commands

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
	trelloService, err := getTrelloService()
	if err != nil {
		return err
	}

	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		err = trelloService.GetShoppingList()
		if err != nil {
			return err
		}

		return nil
	}

	if commandParams[1] == "add" {
		err = trelloService.AddItemToShoppingList(commandParams[2])
		if err != nil {
			return err
		}
	}

	if commandParams[1] == "delete" || commandParams[1] == "del" {
		err = trelloService.DeleteItemFromShoppingList(commandParams[2])
		if err != nil {
			return err
		}
	}

	return nil
}
