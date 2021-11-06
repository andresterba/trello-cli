package commands

import (
	"fmt"
)

type todoCommand struct {
}

func init() {
	RegisterCommand(todoCommand{})
}

func (command todoCommand) GetInformation() (string, string) {
	return "todo", "Print current todo's."
}

func (command todoCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "todo" || commandParams[0] == "t" {
		return true
	}

	return false
}

func (command todoCommand) Execute(commandParams []string) error {
	todoService, context, err := getTodoService()
	if err != nil {
		return err
	}

	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		fmt.Printf("Tasks that are due %s for context %s:\n", red("today"), red(context))
		err = todoService.GetCardsThatAreDueToday()
		if err != nil {
			return err
		}

		return nil
	}

	switch commandParams[1] {
	case "month":
		fmt.Printf("Tasks that are due %s for context %s:\n", red("this month"), red(context))
		err = todoService.GetCardsThatAreDueThisMonth()
		if err != nil {
			return err
		}
	case "overdue":
		fmt.Printf("Tasks that are %s for context %s:\n", red("overdue"), red(context))
		err = todoService.GetCardsThatAreOverDue()
		if err != nil {
			return err
		}
	}

	return nil
}
