package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func (ts *TrelloService) GetShoppingList() error {
	checklist, err := ts.getShoppingCardChecklist()
	if err != nil {
		return err
	}

	printCheckListItems(checklist)

	return nil
}

func (ts *TrelloService) AddItemToShoppingList(itemName string) error {
	checklist, err := ts.getShoppingCardChecklist()
	if err != nil {
		return err
	}

	_, err = ts.client.CreateCheckItem(checklist, itemName, trello.Defaults())
	if err != nil {
		return err
	}

	return nil
}

func (ts *TrelloService) DeleteItemFromShoppingList(itemName string) error {
	checklist, err := ts.getShoppingCardChecklist()
	if err != nil {
		return err
	}

	for _, checkItem := range checklist.CheckItems {
		if checkItem.Name == itemName {
			err = ts.client.Delete(
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

func (ts *TrelloService) getShoppingCardChecklist() (*trello.Checklist, error) {
	board, err := ts.client.GetBoard(ts.config.ShoppingBoardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("could not find board with ID %s", ts.config.ShoppingBoardID)
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
