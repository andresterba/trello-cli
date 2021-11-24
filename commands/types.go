package commands

type Command interface {
	GetInformation() (string, string)
	IsForCommand([]string) bool
	Execute([]string) error
	registerSubCommands()
}

type (
	subCommandFunction func(params []string) error
)

var registeredCommands []Command

func RegisterCommand(command Command) {
	registeredCommands = append(registeredCommands, command)
}

func GetRegisteredCommands() []Command {
	return registeredCommands
}
