package commands

import "errors"

type shoppingListCommand struct {
	subCommands map[string]subCommandFunction
}

func init() {
	slc := shoppingListCommand{}
	slc.subCommands = make(map[string]subCommandFunction)
	slc.registerSubCommands()

	RegisterCommand(slc)
}

func (command shoppingListCommand) GetInformation() (string, string) {
	return "shopping-list", "Interact with your shopping-list."
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

	possibleSubCommand := commandParams[1]
	subCommandFn := command.subCommands[possibleSubCommand]

	if subCommandFn == nil {
		return errors.New("could not find command")
	}

	err = subCommandFn(commandParams[1:])
	if err != nil {
		return err
	}

	return nil
}

func (command shoppingListCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command shoppingListCommand) registerSubCommands() {
	command.registerSubCommand("add", command.SubCommandAdd)
	command.registerSubCommand("delete", command.SubCommandAdd)
	command.registerSubCommand("del", command.SubCommandAdd)
}

func (command shoppingListCommand) SubCommandAdd(commandParams []string) error {
	shoppingListService, err := getShoppingListService()
	if err != nil {
		return err
	}

	err = shoppingListService.AddItemToShoppingList(commandParams[1])
	if err != nil {
		return err
	}

	return nil
}

func (command shoppingListCommand) SubCommandDelete(commandParams []string) error {
	shoppingListService, err := getShoppingListService()
	if err != nil {
		return err
	}

	err = shoppingListService.DeleteItemFromShoppingList(commandParams[1])
	if err != nil {
		return err
	}

	return nil
}
