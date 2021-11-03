package services

import (
	t "github.com/andresterba/trello-cli/trello"
)

type ShoppingListService struct {
	trelloService        *t.TrelloService
	shoppingListBoardID  string
	shoppingListCardName string
}

func NewShoppingListService(
	trelloService *t.TrelloService,
	shoppingListBoardID, shoppingListCardName string,
) *ShoppingListService {
	return &ShoppingListService{
		trelloService,
		shoppingListBoardID,
		shoppingListCardName,
	}
}

func (sls *ShoppingListService) GetShoppingList() error {
	checklist, err := sls.trelloService.GetChecklistFromCard(
		sls.shoppingListBoardID,
		sls.shoppingListCardName,
	)
	if err != nil {
		return err
	}

	t.PrintCheckListItems(checklist)

	return nil
}

func (sls *ShoppingListService) AddItemToShoppingList(itemName string) error {
	checklist, err := sls.trelloService.GetChecklistFromCard(
		sls.shoppingListBoardID,
		sls.shoppingListCardName,
	)
	if err != nil {
		return err
	}

	err = sls.trelloService.CreateChecklistItem(checklist, itemName)
	if err != nil {
		return err
	}

	return nil
}

func (sls *ShoppingListService) DeleteItemFromShoppingList(itemName string) error {
	checklist, err := sls.trelloService.GetChecklistFromCard(
		sls.shoppingListBoardID,
		sls.shoppingListCardName,
	)
	if err != nil {
		return err
	}

	err = sls.trelloService.DeleteChecklistItem(checklist, itemName)
	if err != nil {
		return err
	}

	return nil
}
