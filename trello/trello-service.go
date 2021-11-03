package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/andresterba/trello-cli/config"
)

type TrelloService struct {
	config *config.Config
	client *trello.Client
}

func CreateNewTrelloService(config *config.Config) (*TrelloService, error) {
	client := trello.NewClient(config.AppKey, config.Token)
	if !isTrelloClientWorking(client, config.PersonalConfig.BoardID) {
		return nil, fmt.Errorf(
			"could not connect to the trello api. Please check your tokens or the board id",
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

func (ts *TrelloService) GetAllCardsOnBoard(boardID string) ([]*trello.Card, error) {
	board, err := ts.client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("could not find board with ID %s", boardID)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (ts *TrelloService) GetChecklistFromCard(boardID string, cardName string) (*trello.Checklist, error) {
	board, err := ts.client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("could not find board with ID %s", boardID)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.Name == cardName {
			checklist, err := ts.client.GetChecklist(card.IDCheckLists[0], trello.Defaults())
			if err != nil {
				return nil, err
			}

			return checklist, nil
		}
	}

	return nil, fmt.Errorf(
		"could not find shopping list checklist with name %s",
		ts.config.PersonalConfig.BoardID,
	)
}

func (ts *TrelloService) DeleteChecklistItem(checklist *trello.Checklist, checkItemName string) error {
	for _, checkItem := range checklist.CheckItems {
		if checkItem.Name == checkItemName {
			err := ts.client.Delete(
				fmt.Sprintf("checklists/%s/checkItems/%s", checklist.ID, checkItem.ID),
				trello.Defaults(),
				checklist,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ts *TrelloService) CreateChecklistItem(checklist *trello.Checklist, checkItemName string) error {
	_, err := ts.client.CreateCheckItem(checklist, checkItemName, trello.Defaults())
	if err != nil {
		return err
	}

	return nil
}
