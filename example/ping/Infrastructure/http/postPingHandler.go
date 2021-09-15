package http

import (
	"encoding/json"
	"net/http"

	"example/ping/Application"

	"github.com/vicent-dev/ddd-go/cqrs"
	"github.com/vicent-dev/ddd-go/http/response"
)

type PostPingHandler struct {
	cb cqrs.CommandBus
}

type PostPingHandlerResponse struct{}

func (h PostPingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.cb.Handle(Application.PingCommand{})

	if err != nil {
		response.WriteErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(PostPingHandlerResponse{})
}
