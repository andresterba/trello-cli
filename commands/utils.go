package commands

import (
	"errors"

	"github.com/andresterba/trello-cli/config"
	"github.com/andresterba/trello-cli/services"
	"github.com/andresterba/trello-cli/trello"
)

func getConfig() (*config.Config, error) {
	config, err := config.LoadConfig(config.GetConfigPath())
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getTrelloService() (*trello.TrelloService, error) {
	config, err := getConfig()
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

	trelloBoardID, err := getBoardIDForCurrentContext()
	if err != nil {
		return nil, err
	}

	return services.NewTodoService(
		trelloService,
		trelloBoardID,
	), nil
}

func getShoppingListService() (*services.ShoppingListService, error) {
	trelloService, err := getTrelloService()
	if err != nil {
		return nil, err
	}
	config, err := getConfig()
	if err != nil {
		return nil, err
	}

	return services.NewShoppingListService(
		trelloService,
		config.ShoppingConfig.BoardID,
		config.ShoppingConfig.ListCardName,
	), nil
}

func getBoardIDForCurrentContext() (string, error) {
	config, err := getConfig()
	if err != nil {
		return "", err
	}

	switch config.DefaultContext {
	case PersonalContext:
		return config.PersonalConfig.BoardID, nil
	case WorkContext:
		return config.WorkConfig.BoardID, nil
	}

	return "", errors.New("default context in config is not valid")
}

func getCurrentContext() (string, error) {
	config, err := getConfig()
	if err != nil {
		return "", err
	}

	return config.DefaultContext, nil
}
