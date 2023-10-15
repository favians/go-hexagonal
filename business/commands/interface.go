package commands

type Service interface {
	ProcessCommand(commandSpec CommandSpec) error
}