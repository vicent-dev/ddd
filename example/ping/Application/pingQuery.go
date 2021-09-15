package Application

import (
	"example/ping/Domain"

	"github.com/vicent-dev/ddd-go/cqrs"
)

type PingQuery struct{}

func (PingQuery PingQuery) ID() string {
	return "ping.query"
}

type PingQueryResult struct {
	Value map[string]string
}

type PingQueryHandler struct {
	Pr Domain.PingRepository
}

func (handler PingQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	ping := handler.Pr.Get()
	return PingQueryResult{map[string]string{"ping": ping.Value}}, nil
}
