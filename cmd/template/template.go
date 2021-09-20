package template

const (
	domainTemplate = `package Domain

type %v struct {}
`

	queryTempalte = `package Application

import (
	"github.com/vicent-dev/ddd"
)

type %vQuery struct{}

func (%v %vQuery) ID() string {
	return "%v"
}

type %vResult struct { }

type %vHandler struct { }

func (handler %vHandler) Handle(query ddd.Query) (interface{}, error) {
	return %vResult{}, nil
}
	`

	commandTempalte = `package Application

import (
	"github.com/vicent-dev/ddd"
)

type %vCommand struct{}

func (%v %vCommand) ID() string {
	return "%v"
}

type %vHandler struct { }

func (handler %vHandler) Handle(command ddd.Command) error {
	return nil
}
	`

	serviceProviderTemplate = `package Infrastructure

import (
	"github.com/vicent-dev/ddd"
)

func NewServiceProvider() *ddd.ServiceProvider {
	sp := ddd.NewServiceProvider()


	//query bus
	queryBus := ddd.NewQueryBus()

	//QUERY_SERVICES
	//QUERY_SERVICES

	sp.Register("query_bus", queryBus)

	//command bus
	commandBus := ddd.NewCommandBus()


	//COMMAND_SERVICES
	//COMMAND_SERVICES


	sp.Register("command_bus", commandBus)

	return sp
}`
)
