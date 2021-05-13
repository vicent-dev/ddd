package http

import (
	"ddd-go/cqrs"
	"ddd-go/example/ping/Application"
	"ddd-go/http/response"
	"encoding/json"
	"net/http"
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
