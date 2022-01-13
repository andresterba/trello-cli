package commands

import (
	"errors"
	"fmt"
)

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
	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		fmt.Println("Use 'boards' or 'cards' to get specific IDs.")

		return nil
	}

	possibleSubCommand := commandParams[1]
	subCommandFn := command.subCommands[possibleSubCommand]

	if subCommandFn == nil {
		return errors.New("could not find command")
	}

	err := subCommandFn(commandParams[1:])
	if err != nil {
		return err
	}

	return nil
}

func (command listCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command listCommand) registerSubCommands() {
	command.registerSubCommand("boards", command.listAllBoards)
	command.registerSubCommand("cards", command.listAllCardsOnBoard)
	command.registerSubCommand("lists", command.listAllListsOnBoard)
}

func (command listCommand) listAllBoards(params []string) error {
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

func (command listCommand) listAllCardsOnBoard(params []string) error {
	if len(params) != 2 {
		fmt.Println("Please provide a single board ID.")

		return nil
	}

	trelloService, err := getTrelloService()
	if err != nil {
		return err
	}

	cards, err := trelloService.GetAllCardsOnBoard(params[1])
	if err != nil {
		return err
	}

	fmt.Printf("%-24s: %-24s\n", "card.ID", "card.Name")
	fmt.Printf("%-24s: %-24s\n", "-----------------------", "-----------------------")
	for _, card := range cards {
		fmt.Printf("%-24s: %-24s\n", card.ID, card.Name)
	}

	return nil
}

func (command listCommand) listAllListsOnBoard(params []string) error {
	if len(params) != 2 {
		fmt.Println("Please provide a single board ID.")

		return nil
	}

	trelloService, err := getTrelloService()
	if err != nil {
		return err
	}

	lists, err := trelloService.GetAllListsOnBoard(params[1])
	if err != nil {
		return err
	}

	fmt.Printf("%-24s: %-24s\n", "list.ID", "list.Name")
	fmt.Printf("%-24s: %-24s\n", "---------------------------------------", "-----------------------")
	for _, list := range lists {
		fmt.Printf("%-24s: %-24s\n", list.ID, list.Name)
	}

	return nil
}
