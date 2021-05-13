package Application

import (
	"ddd-go/cqrs"
	"ddd-go/example/ping/Domain"
)

type PingCommand struct{}

func (p PingCommand) ID() string {
	return "ping.command"
}

type PingCommandHandler struct {
	Pr Domain.PingRepository
}

func (handler PingCommandHandler) Handle(command cqrs.Command) error {
	//use data from command in real world use case
	err := handler.Pr.Create(Domain.CreatePing())
	return err
}