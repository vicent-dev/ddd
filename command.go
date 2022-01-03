package ddd

import "errors"

type CommandHandler[T comparable]interface {
	Handle(command T) error
}

type CommandBus[T comparable] struct {
	handlers map[T]CommandHandler[T]
}

func NewCommandBus[T comparable]() CommandBus[T] {
	return CommandBus[T]{make(map[T]CommandHandler[T])}
}

func (b *CommandBus[T]) AddHandler(command T, handler CommandHandler[T]) error {

	if _, added := b.handlers[command]; added {
		return errors.New("The bus already has a handler associated to that command id")
	}

	b.handlers[command] = handler

	return nil
}

func (b *CommandBus[T]) Handle(command T) error {
	if _, added := b.handlers[command]; !added {
		return errors.New("Bus doesn't have a valid handler for that command")
	}

	return b.handlers[command].Handle(command)
}
