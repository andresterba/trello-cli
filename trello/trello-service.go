package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"gitlab.cloudf.de/andre/trello-cli/config"
)

type TrelloService struct {
	config *config.Config
	client *trello.Client
}

func CreateNewTrelloService(config *config.Config) (*TrelloService, error) {
	client := trello.NewClient(config.AppKey, config.Token)
	if !isTrelloClientWorking(client, config.BoardID) {
		return nil, fmt.Errorf(
			"could not connect to the trello api. Please check your tokens or the board  id",
		)
	}

	return &TrelloService{
		config: config,
		client: client,
	}, nil
}

func isTrelloClientWorking(client *trello.Client, board string) bool {
	_, err := client.GetBoard(board, trello.Defaults())

	return err == nil
}
