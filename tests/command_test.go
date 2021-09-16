package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type TestCommand struct{}

func (t TestCommand) ID() string {
	return "test_command"
}

type TestCommandHandler struct{}

func (t TestCommandHandler) Handle(command ddd.Command) error {
	return nil
}

func CommandBus_AddHandlerAndHandleTest(t *testing.T) {
	cb := ddd.NewCommandBus()

	if err := cb.AddHandler(TestCommand{}, TestCommandHandler{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

	if err := cb.Handle(TestCommand{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

}
