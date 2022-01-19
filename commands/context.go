package commands

type contextCommand struct {
	subCommands map[string]subCommandFunction
}

const (
	WorkContext     = "work"
	PersonalContext = "personal"
	ProjectsContext = "projects"
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
	newContext := commandParams[1]

	err := changeContextTo(newContext)
	if err != nil {
		return err
	}

	return nil
}

func changeContextTo(couldBeNewContext string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	if checkIfContextIsValid(couldBeNewContext) {
		config.DefaultContext = couldBeNewContext
	}

	err = config.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func checkIfContextIsValid(context string) bool {
	if context == PersonalContext || context == WorkContext || context == ProjectsContext {
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
