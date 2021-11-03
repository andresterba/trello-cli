package commands

import "fmt"

type todoCommand struct {
}

func init() {
	RegisterCommand(todoCommand{})
}

func (command todoCommand) GetInformation() (string, string) {
	return "version", "Print current version."
}

func (command todoCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "todo" || commandParams[0] == "t" {
		return true
	}

	return false
}

func (command todoCommand) Execute(commandParams []string) error {
	todoService, err := getTodoService()
	if err != nil {
		return err
	}

	commandParamsLength := len(commandParams)

	if commandParamsLength == 1 {
		fmt.Println("Tasks that are due today:")
		err = todoService.GetCardsThatAreDueToday()
		if err != nil {
			return err
		}

		return nil
	}

	if commandParams[1] == "month" {
		err = todoService.GetCardsThatAreDueThisMonth()
		if err != nil {
			return err
		}
	}

	return nil
}
