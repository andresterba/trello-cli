package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func PrintCheckListItems(checklist *trello.Checklist) {
	for _, item := range checklist.CheckItems {
		printCheckListItem(item)
	}
}

func printCheckListItem(item trello.CheckItem) {
	if isCheckItemCompleted(item) {
		printCompletedChecklistItem(item)

		return
	}

	printUncompletedChecklistItem(item)

}

func isCheckItemCompleted(item trello.CheckItem) bool {
	return item.State != "incomplete"
}

func printCompletedChecklistItem(item trello.CheckItem) {
	fmt.Printf("* [x] %s\n", item.Name)
}

func printUncompletedChecklistItem(item trello.CheckItem) {
	fmt.Printf("* [ ] %s\n", item.Name)
}
