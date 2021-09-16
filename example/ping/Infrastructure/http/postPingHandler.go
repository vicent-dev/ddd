package http

import (
	"encoding/json"
	"net/http"

	"example/ping/Application"

	"github.com/vicent-dev/ddd"
)

type PostPingHandler struct {
	cb ddd.CommandBus
}

type PostPingHandlerResponse struct{}

func (h PostPingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.cb.Handle(Application.PingCommand{})

	if err != nil {
		ddd.WriteErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(PostPingHandlerResponse{})
}
