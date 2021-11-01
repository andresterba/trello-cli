package commands

import (
	"gitlab.cloudf.de/andre/trello-cli/config"
	"gitlab.cloudf.de/andre/trello-cli/trello"
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
