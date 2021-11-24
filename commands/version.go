package commands

import "fmt"

type versionCommand struct {
}

func init() {
	RegisterCommand(versionCommand{})
}

func (command versionCommand) GetInformation() (string, string) {
	return "shopping-list", "Interact with our shopping-list."
}

func (command versionCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "version" || commandParams[0] == "--version" {
		return true
	}

	return false
}

func (command versionCommand) Execute(commandParams []string) error {
	fmt.Println("0.0.0")
	return nil
}

func (command versionCommand) registerSubCommands() {
}
