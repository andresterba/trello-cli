package commands

type contextCommand struct {
}

const (
	WorkContext     = "work"
	PersonalContext = "personal"
)

func init() {
	RegisterCommand(contextCommand{})
}

func (command contextCommand) GetInformation() (string, string) {
	return "context", "Set context for other commands."
}

func (command contextCommand) IsForCommand(commandParams []string) bool {
	if commandParams[0] == "context" || commandParams[0] == "c" {
		return true
	}

	return false
}

func (command contextCommand) Execute(commandParams []string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	newContext := commandParams[1]

	if checkIfContextIsValid(newContext) {
		config.DefaultContext = newContext
	}

	err = config.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func checkIfContextIsValid(context string) bool {
	if context == PersonalContext || context == WorkContext {
		return true
	}

	return false
}
