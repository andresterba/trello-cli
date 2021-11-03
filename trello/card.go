package trello

import (
	"fmt"
	"sort"

	"github.com/adlio/trello"
)

func printCard(position int, card *trello.Card) {
	fmt.Printf("%d. %s Due: %s\n", position, card.Name, card.Due.Format("Mon _2 Jan"))
}

func sortCardsByDueDate(cards []*trello.Card) []*trello.Card {
	sort.Slice(
		cards,
		func(i, j int) bool {
			return cards[i].Due.Before(*cards[j].Due)
		},
	)

	return cards
}

func printCards(cards []*trello.Card) {
	for position, card := range cards {
		printCard(position+1, card)
	}
}
