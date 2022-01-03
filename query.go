package ddd

import "errors"

type Query interface {
	ID() string
}

type QueryHandler[T comparable, K any] interface {
	Handle(query T) (K, error)
}

type QueryBus[T comparable, K any] struct {
	handlers map[T]QueryHandler[T, K]
}

func NewQueryBus[T comparable, K any]() QueryBus[T, K] {
	return QueryBus[T, K]{make(map[T]QueryHandler[T, K])}
}

func (b QueryBus[T, K]) AddHandler(query T, handler QueryHandler[T, K]) error {

	if _, added := b.handlers[query]; added {
		return errors.New("The bus already has a handler associated to that query id")
	}

	b.handlers[query] = handler

	return nil
}

func (b *QueryBus[T, K]) Handle(query T) (interface{}, error) {
	if _, added := b.handlers[query]; !added {
		return nil, errors.New("Bus doesn't have a valid handler for that query")
	}

	return b.handlers[query].Handle(query)
}
