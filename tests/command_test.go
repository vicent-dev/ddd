package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type TestCommand struct{}

type TestCommandHandler[T comparable] struct{}

func (t TestCommandHandler[T]) Handle(command T) error {
	return nil
}

func CommandBus_AddHandlerAndHandleTest(t *testing.T) {
	cb := ddd.NewCommandBus[TestCommand]()

	if err := cb.AddHandler(TestCommand{}, TestCommandHandler[TestCommand]{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

	if err := cb.Handle(TestCommand{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

}
