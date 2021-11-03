package trello

import (
	"fmt"
	"time"

	"github.com/adlio/trello"
)

func (ts *TrelloService) GetCardsThatAreDueToday() error {
	cards, err := ts.getAllCardsOnBoard()
	if err != nil {
		return err
	}

	var cardsDueToday []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueToday(*card.Due) {
			cardsDueToday = append(cardsDueToday, card)
		}
	}

	cardsDueToday = sortCardsByDueDate(cardsDueToday)

	printCards(cardsDueToday)

	return nil
}

func (ts *TrelloService) GetCardsThatAreDueThisMonth() error {
	cards, err := ts.getAllCardsOnBoard()
	if err != nil {
		return err
	}

	var cardsDueThisMonth []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueThisMonth(*card.Due) {
			cardsDueThisMonth = append(cardsDueThisMonth, card)
		}
	}

	cardsDueThisMonth = sortCardsByDueDate(cardsDueThisMonth)

	printCards(cardsDueThisMonth)

	return nil
}

func (ts *TrelloService) getAllCardsOnBoard() ([]*trello.Card, error) {
	board, err := ts.client.GetBoard(ts.config.TodoBoardID, trello.Defaults())
	if err != nil {
		return nil, fmt.Errorf("could not find board with ID %s", ts.config.TodoBoardID)
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func isDueSetOnCard(card *trello.Card) bool {
	return card.Due != nil
}

func isCardDueToday(dueTime time.Time) bool {
	year, month, day := dueTime.Date()
	yearNow, monthNow, dayNow := time.Now().Date()

	if (year == yearNow) && (month == monthNow) && (day == dayNow) {
		return true
	}

	return false
}

func isCardDueThisMonth(dueTime time.Time) bool {
	year, month, _ := dueTime.Date()
	yearNow, monthNow, _ := time.Now().Date()

	if (year == yearNow) && (month == monthNow) {
		return true
	}

	return false
}
