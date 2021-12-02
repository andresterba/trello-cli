package commands

type contextCommand struct {
	subCommands map[string]subCommandFunction
}

const (
	WorkContext     = "work"
	PersonalContext = "personal"
)

func init() {
	cc := contextCommand{}
	cc.subCommands = make(map[string]subCommandFunction)

	RegisterCommand(cc)
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

//lint:ignore U1000 will be implemented soon
func (command contextCommand) registerSubCommand(name string, fn subCommandFunction) {
	command.subCommands[name] = fn
}

func (command contextCommand) registerSubCommands() {
}
