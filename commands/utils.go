package commands

import (
	"github.com/andresterba/trello-cli/config"
	"github.com/andresterba/trello-cli/services"
	"github.com/andresterba/trello-cli/trello"
)

func getTrelloService() (*trello.TrelloService, error) {
	config, err := config.LoadConfig(config.GetConfigPath())
	if err != nil {
		return nil, err
	}
	trelloService, err := trello.CreateNewTrelloService(
		config,
	)
	if err != nil {
		return nil, err
	}

	return trelloService, nil
}

func getTodoService() (*services.TodoService, error) {
	trelloService, err := getTrelloService()
	if err != nil {
		return nil, err
	}
	config, err := config.LoadConfig(config.GetConfigPath())
	if err != nil {
		return nil, err
	}

	return services.NewTodoService(
		trelloService,
		config.TodoBoardID,
	), nil
}

func getShoppingListService() (*services.ShoppingListService, error) {
	trelloService, err := getTrelloService()
	if err != nil {
		return nil, err
	}
	config, err := config.LoadConfig(config.GetConfigPath())
	if err != nil {
		return nil, err
	}

	return services.NewShoppingListService(
		trelloService,
		config.ShoppingBoardID,
		config.ShoppingListCardName,
	), nil
}
