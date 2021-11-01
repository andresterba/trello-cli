package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func printCheckListItems(checklist *trello.Checklist) {
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
	if item.State == "incomplete" {
		return false
	}

	return true
}

func printCompletedChecklistItem(item trello.CheckItem) {
	fmt.Printf("* [x] %s\n", item.Name)
}

func printUncompletedChecklistItem(item trello.CheckItem) {
	fmt.Printf("* [ ] %s\n", item.Name)
}
