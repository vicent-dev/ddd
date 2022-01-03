package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type TestQuery struct { }

//type T and K??
type TestQueryHandler[T comparable, K any] struct{}

func (t TestQueryHandler[T, K]) Handle(query T) (K, error) {
	var res K
	return res, nil
}

func QueryBus_AddHandlerAndHandleTest(t *testing.T) {
	qb := ddd.NewQueryBus[any, any]()
	qh := TestQueryHandler[any, any]{}

	q := TestQuery{}

	if err := qb.AddHandler(q, qh); err != nil { t.Errorf("Not expected error, get: %s", err.Error())
	}

	if _, err := qb.Handle(q); err != nil {
		t.Errorf("Not expected error, get: %s", err.Error())
	}

}
