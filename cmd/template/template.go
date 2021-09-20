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
)
