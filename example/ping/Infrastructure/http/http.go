package http

import (
	"github.com/gorilla/mux"
	"github.com/vicent-dev/ddd"
)

type PingHandler struct {
	getPingHandler  GetPingHandler
	postPingHandler PostPingHandler
}

func NewHandler(sp *ddd.ServiceProvider) *PingHandler {
	return &PingHandler{
		GetPingHandler{sp.Get("query_bus").(ddd.QueryBus)},
		PostPingHandler{sp.Get("command_bus").(ddd.CommandBus)},
	}
}

func (ph *PingHandler) SetEndpoints(r *mux.Router) {
	r.HandleFunc("/ping", ph.getPingHandler.Handle).Methods("GET")
	r.HandleFunc("/ping", ph.postPingHandler.Handle).Methods("POST")
}
