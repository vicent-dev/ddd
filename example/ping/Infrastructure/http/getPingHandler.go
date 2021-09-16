package http

import (
	"encoding/json"
	"net/http"

	"example/ping/Application"

	"github.com/vicent-dev/ddd"
)

type GetPingHandler struct {
	qb ddd.QueryBus
}

type GetPingHandlerResponse struct {
	//add more stuff if needed
	PingResult interface{}
}

func (h GetPingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	result, err := h.qb.Handle(Application.PingQuery{})

	if err != nil {
		ddd.WriteErrorResponse(w, err)
		return
	}

	json.NewEncoder(w).Encode(GetPingHandlerResponse{result})
}
