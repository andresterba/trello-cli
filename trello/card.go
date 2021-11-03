package trello

import (
	"fmt"
	"sort"

	"github.com/adlio/trello"
)

func PrintCard(position int, card *trello.Card) {
	fmt.Printf("%d. %s Due: %s\n", position, card.Name, card.Due.Format("Mon _2 Jan"))
}

func SortCardsByDueDate(cards []*trello.Card) []*trello.Card {
	sort.Slice(
		cards,
		func(i, j int) bool {
			return cards[i].Due.Before(*cards[j].Due)
		},
	)

	return cards
}

func PrintCards(cards []*trello.Card) {
	for position, card := range cards {
		PrintCard(position+1, card)
	}
}
