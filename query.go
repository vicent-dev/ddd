package ddd

import "errors"

type Query interface {
	ID() string
}

type QueryHandler interface {
	Handle(Query Query) (interface{}, error)
}

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() QueryBus {
	return QueryBus{make(map[string]QueryHandler)}
}

func (b *QueryBus) AddHandler(Query Query, handler QueryHandler) error {

	if _, added := b.handlers[Query.ID()]; added {
		return errors.New("The bus already has a handler associated to that query id")
	}

	b.handlers[Query.ID()] = handler

	return nil
}

func (b *QueryBus) Handle(query Query) (interface{}, error) {
	if _, added := b.handlers[query.ID()]; !added {
		return nil, errors.New("Bus doesn't have a valid handler for that query")
	}

	return b.handlers[query.ID()].Handle(query)
}
