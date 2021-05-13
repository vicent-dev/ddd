package http

import (
	"ddd-go/cqrs"
	"ddd-go/example/ping/Application"
	"ddd-go/http/response"
	"encoding/json"
	"net/http"
)

type GetPingHandler struct {
	qb cqrs.QueryBus
}

type GetPingHandlerResponse struct {
	//add more stuff if needed
	pingResult interface{}
}

func (h GetPingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	result, err := h.qb.Handle(Application.PingQuery{})

	if err != nil {
		response.WriteErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(GetPingHandlerResponse{result})
}
