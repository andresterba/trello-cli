package commands

import (
	"errors"
	"fmt"
)

type todoCommand struct {
	subCommands map[string]subCommandFunction
}

func init() {
	tc := todoCommand{}
	tc.subCommands = make(map[string]subCommandFunction)
	tc.registerSubCommands()

	RegisterCommand(tc)
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
	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		err := command.getOverallTasksDueThisWeek()
		if err != nil {
			return err
		}

		return nil
	}

	possibleSubCommand := commandParams[1]
	subCommandFn := command.subCommands[possibleSubCommand]

	if subCommandFn == nil {
		return errors.New("could not find command")
	}

	err := subCommandFn(commandParams[1:])
	if err != nil {
		return err
	}

	return nil
}

func (command todoCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command todoCommand) registerSubCommands() {
	command.registerSubCommand("month", command.subCommandDueThisMonth)
	command.registerSubCommand("week", command.subCommandDueThisWeek)
	command.registerSubCommand("overdue", command.subCommandOverdue)
}

func (command todoCommand) subCommandDueThisMonth(commandParams []string) error {
	todoService, context, err := getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("Tasks that are due %s for context %s:\n", red("this month"), red(context))
	err = todoService.GetCardsThatAreDueThisMonth()
	if err != nil {
		return err
	}

	return nil
}

func (command todoCommand) subCommandOverdue(commandParams []string) error {
	todoService, context, err := getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("Tasks that are %s for context %s:\n", red("overdue"), red(context))
	err = todoService.GetCardsThatAreOverDue()
	if err != nil {
		return err
	}

	return nil
}

func (command todoCommand) subCommandDueThisWeek(commandParams []string) error {
	todoService, context, err := getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("Tasks that are due %s for context %s:\n", red("this week"), red(context))
	err = todoService.GetCardsThatAreDueThisWeek()
	if err != nil {
		return err
	}

	return nil
}

func (command todoCommand) getOverallTasksDueThisWeek() error {
	changeContextTo(PersonalContext)
	todoService, context, err := getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("Tasks that are due %s for context %s:\n", red("this week"), red(context))
	err = todoService.GetCardsThatAreDueThisWeek()
	if err != nil {
		return err
	}

	changeContextTo(WorkContext)
	todoService, context, err = getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("\nTasks that are due %s for context %s:\n", red("this week"), red(context))
	err = todoService.GetCardsThatAreDueThisWeek()
	if err != nil {
		return err
	}

	changeContextTo(ProjectsContext)
	todoService, context, err = getTodoService()
	if err != nil {
		return err
	}

	fmt.Printf("\nTasks that are due %s for context %s:\n", red("this week"), red(context))
	err = todoService.GetCardsThatAreDueThisWeek()
	if err != nil {
		return err
	}

	return nil
}
