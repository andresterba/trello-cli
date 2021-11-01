package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func (ts *TrelloService) getShoppingCardChecklist() (*trello.Checklist, error) {
	board, err := ts.client.GetBoard(ts.config.BoardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("could not find board with ID %s", ts.config.BoardID)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.Name == ts.config.ShoppingListCardName {
			checklist, err := ts.client.GetChecklist(card.IDCheckLists[0], trello.Defaults())
			if err != nil {
				return nil, err
			}

			return checklist, nil
		}
	}

	return nil, fmt.Errorf(
		"could not find shopping list checklist with name %s",
		ts.config.ShoppingListCardName,
	)
}
