package trello

import (
	"github.com/adlio/trello"
	"gitlab.cloudf.de/andre/trello-cli/config"
)

type TrelloService struct {
	config *config.Config
	client *trello.Client
}

func CreateNewTrelloService(config *config.Config) (error, *TrelloService) {
	client := trello.NewClient(config.AppKey, config.Token)
	// TODO: as client does not return an error, do an api call and check if it does error.

	return nil, &TrelloService{
		config: config,
		client: client,
	}
}

func (ts *TrelloService) GetShoppingList() error {
	board, err := ts.client.GetBoard(ts.config.BoardID, trello.Defaults())
	if err != nil {
		return err
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return err
	}

	for _, card := range cards {
		if card.Name == ts.config.ShoppingListCardName {
			checklist, err := ts.client.GetChecklist(card.IDCheckLists[0], trello.Defaults())
			if err != nil {
				return err
			}
			printCheckListItems(checklist)
		}
	}

	return nil
}
