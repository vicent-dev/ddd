package http

import (
	"ddd-go/cqrs"
	"net/http"
)

type GetPingHandler struct {
	qb cqrs.QueryBus
}

func (h GetPingHandler) Handle(w http.ResponseWriter, r *http.Request) {
}
