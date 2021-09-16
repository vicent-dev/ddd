package Application

import (
	"example/ping/Domain"

	"github.com/vicent-dev/ddd"
)

type PingCommand struct{}

func (p PingCommand) ID() string {
	return "ping.command"
}

type PingCommandHandler struct {
	Pr Domain.PingRepository
}

func (handler PingCommandHandler) Handle(command ddd.Command) error {
	//use data from command in real world use case
	err := handler.Pr.Create(Domain.CreatePing())
	return err
}
