package commands

import (
	"errors"
	"fmt"

	"github.com/adlio/trello"
)

type recurringCommand struct {
	subCommands map[string]subCommandFunction
}

func init() {
	tc := recurringCommand{}
	tc.subCommands = make(map[string]subCommandFunction)
	tc.registerSubCommands()

	RegisterCommand(tc)
}

func (command recurringCommand) GetInformation() (string, string) {
	return "recurring", "Print or add recurring todo's."
}

func (command recurringCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "recurring" || commandParams[0] == "r" {
		return true
	}

	return false
}

func (command recurringCommand) Execute(commandParams []string) error {
	err := command.setDefaultContextToPersonal()
	if err != nil {
		return err
	}

	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		return errors.New("tbd")
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

func (command recurringCommand) setDefaultContextToPersonal() error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	config.DefaultContext = "personal"

	err = config.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func (command recurringCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command recurringCommand) registerSubCommands() {
	command.registerSubCommand("add", command.subCommandAdd)
}

func (command recurringCommand) subCommandAdd(commandParams []string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	todoService, _, err := getTodoService()
	if err != nil {
		return err
	}

	trelloService, err := getTrelloService()
	if err != nil {
		return err
	}

	existingCards, err := trelloService.GetAllCardsOnBoard(config.PersonalConfig.BoardID)
	if err != nil {
		return err
	}
	existingCardsLookupMap := make(map[string]*trello.Card)

	for _, card := range existingCards {
		existingCardsLookupMap[card.Name] = card
	}

	for _, rt := range config.PersonalConfig.RecurringTasks {
		_, alreadyExists := existingCardsLookupMap[rt.Name]
		if !alreadyExists {
			err := todoService.CreateNewCard(rt.Name, rt.ListID)
			if err != nil {
				return err
			}

			fmt.Printf("Added recurring task '%s'\n", rt.Name)
		}
	}

	return nil
}
