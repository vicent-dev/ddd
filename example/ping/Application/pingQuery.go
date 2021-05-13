package Application

import (
	"ddd-go/cqrs"
	"ddd-go/example/ping/Domain"
)

type PingQuery struct{}

func (PingQuery PingQuery) ID() string {
	return "ping.query"
}

type PingQueryResult struct {
	value map[string]string
}

type PingQueryHandler struct {
	Pr Domain.PingRepository
}

func (handler PingQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	ping := handler.Pr.Get()
	return PingQueryResult{map[string]string{"ping": ping.Value}}, nil
}
