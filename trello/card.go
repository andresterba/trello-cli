package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func printCard(card *trello.Card) {
	fmt.Printf("%s\n", card.Name)
}
