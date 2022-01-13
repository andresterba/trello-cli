package services

import (
	"time"

	"github.com/adlio/trello"
	t "github.com/andresterba/trello-cli/trello"
)

type TodoService struct {
	trelloService *t.TrelloService
	todoBoard     string
}

func NewTodoService(
	trelloService *t.TrelloService,
	todoBoard string,
) *TodoService {
	return &TodoService{
		trelloService,
		todoBoard,
	}
}

func (ts *TodoService) CreateNewCard(name string, listID string, labelIDs []string) error {
	card := trello.Card{
		Name:     name,
		IDBoard:  ts.todoBoard,
		IDList:   listID,
		IDLabels: labelIDs,
	}

	err := ts.trelloService.CreateCard(&card)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TodoService) GetCardsThatAreDueToday() error {
	cards, err := ts.trelloService.GetAllCardsOnBoard(ts.todoBoard)
	if err != nil {
		return err
	}

	var cardsDueToday []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueToday(*card.Due) && !isCardDueCompleted(card) {
			cardsDueToday = append(cardsDueToday, card)
		}
	}

	cardsDueToday = t.SortCardsByDueDate(cardsDueToday)

	t.PrintCards(cardsDueToday)

	return nil
}

func (ts *TodoService) GetCardsThatAreDueThisWeek() error {
	cards, err := ts.trelloService.GetAllCardsOnBoard(ts.todoBoard)
	if err != nil {
		return err
	}

	var cardsDueThisWeek []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueThisWeek(*card.Due) && !isCardDueCompleted(card) {
			cardsDueThisWeek = append(cardsDueThisWeek, card)
		}
	}

	cardsDueThisWeek = t.SortCardsByDueDate(cardsDueThisWeek)

	t.PrintCards(cardsDueThisWeek)

	return nil
}

func (ts *TodoService) GetCardsThatAreDueThisMonth() error {
	cards, err := ts.trelloService.GetAllCardsOnBoard(ts.todoBoard)
	if err != nil {
		return err
	}

	var cardsDueThisMonth []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardDueThisMonth(*card.Due) && !isCardDueCompleted(card) {
			cardsDueThisMonth = append(cardsDueThisMonth, card)
		}
	}

	cardsDueThisMonth = t.SortCardsByDueDate(cardsDueThisMonth)

	t.PrintCards(cardsDueThisMonth)

	return nil
}

func (ts *TodoService) GetCardsThatAreOverDue() error {
	cards, err := ts.trelloService.GetAllCardsOnBoard(ts.todoBoard)
	if err != nil {
		return err
	}

	var cardsDueToday []*trello.Card

	for _, card := range cards {
		if isDueSetOnCard(card) && isCardOverDue(*card.Due) && !isCardDueCompleted(card) {
			cardsDueToday = append(cardsDueToday, card)
		}
	}

	cardsDueToday = t.SortCardsByDueDate(cardsDueToday)

	t.PrintCards(cardsDueToday)

	return nil
}

func isDueSetOnCard(card *trello.Card) bool {
	return card.Due != nil
}

func isCardDueCompleted(card *trello.Card) bool {
	return card.DueComplete
}

func isCardDueToday(dueTime time.Time) bool {
	year, month, day := dueTime.Date()
	yearNow, monthNow, dayNow := time.Now().Date()

	if (year == yearNow) && (month == monthNow) && (day == dayNow) {
		return true
	}

	return false
}

func isCardOverDue(dueTime time.Time) bool {
	return dueTime.Before(time.Now())
}

func isCardDueThisMonth(dueTime time.Time) bool {
	year, month, _ := dueTime.Date()
	yearNow, monthNow, _ := time.Now().Date()

	if (year == yearNow) && (month == monthNow) {
		return true
	}

	return false
}

func isCardDueThisWeek(dueTime time.Time) bool {
	yearDue, weekDue := dueTime.ISOWeek()
	yearNow, weekNow := dueTime.ISOWeek()

	if (yearDue == yearNow) && (weekDue == weekNow) {
		return true
	}

	return false
}
