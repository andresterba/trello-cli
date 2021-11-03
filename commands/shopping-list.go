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
	shoppingListService, err := getShoppingListService()
	if err != nil {
		return err
	}

	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		err = shoppingListService.GetShoppingList()
		if err != nil {
			return err
		}

		return nil
	}

	if commandParams[1] == "add" {
		err = shoppingListService.AddItemToShoppingList(commandParams[2])
		if err != nil {
			return err
		}
	}

	if commandParams[1] == "delete" || commandParams[1] == "del" {
		err = shoppingListService.DeleteItemFromShoppingList(commandParams[2])
		if err != nil {
			return err
		}
	}

	return nil
}
