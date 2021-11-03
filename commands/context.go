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
	return nil
}
