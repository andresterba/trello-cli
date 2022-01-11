package commands

import "fmt"

type listCommand struct {
	subCommands map[string]subCommandFunction
}

func init() {
	lc := listCommand{}
	lc.subCommands = make(map[string]subCommandFunction)
	lc.registerSubCommands()

	RegisterCommand(lc)
}

func (command listCommand) GetInformation() (string, string) {
	return "list", "Print board and card ids."
}

func (command listCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "list" || commandParams[0] == "l" {
		return true
	}

	return false
}

func (command listCommand) Execute(commandParams []string) error {
	trelloService, err := getTrelloService()
	if err != nil {
		return err
	}

	boards, err := trelloService.GetAllBoards()
	if err != nil {
		return err
	}

	fmt.Printf("%-20s: %-24s\n", "board.Name", "board.ID")
	fmt.Printf("%-20s: %-24s\n", "-------------------", "-----------------------")
	for _, board := range boards {
		fmt.Printf("%-20s: %-24s\n", board.Name, board.ID)
	}

	return nil
}

func (command listCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command listCommand) registerSubCommands() {
	command.registerSubCommand("boards", command.listAllBoards)
	command.registerSubCommand("cards", command.listAllCardsOnBoard)
}

func (command listCommand) listAllBoards(params []string) error {
	return nil
}

func (command listCommand) listAllCardsOnBoard(params []string) error {
	return nil
}
