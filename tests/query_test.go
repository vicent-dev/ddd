package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type TestQuery struct{}

func (t TestQuery) ID() string {
	return "test_query"
}

type TestQueryHandler struct{}

func (t TestQueryHandler) Handle(query ddd.Query) (interface{}, error) {
	return nil, nil
}

func QueryBus_AddHandlerAndHandleTest(t *testing.T) {
	qb := ddd.NewQueryBus()

	if err := qb.AddHandler(TestQuery{}, TestQueryHandler{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

	if _, err := qb.Handle(TestQuery{}); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

}
