package main

import "errors"

type Command interface {
	ID() string
}

type CommandHandler interface {
	Handle(command Command) error
}

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() CommandBus {
	return CommandBus{make(map[string]CommandHandler)}
}

func (b *CommandBus) AddHandler(command Command, handler CommandHandler) error {

	if _, added := b.handlers[command.ID()]; added {
		return errors.New("The bus already has a handler associated to that command id")
	}

	b.handlers[command.ID()] = handler

	return nil
}

func (b *CommandBus) Handle(command Command) error {
	if _, added := b.handlers[command.ID()]; !added {
		return errors.New("Bus doesn't have a valid handler for that command")
	}

	return b.handlers[command.ID()].Handle(command)
}
