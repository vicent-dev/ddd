package http

import (
	"ddd-go/cqrs"
	"ddd-go/serviceProvider"

	"github.com/gorilla/mux"
)

type PingHandler struct {
	getPingHandler  GetPingHandler
	postPingHandler PostPingHandler
}

func NewHandler(sp *serviceProvider.ServiceProvider) *PingHandler {
	return &PingHandler{
		GetPingHandler{sp.Get("queryBus").(cqrs.QueryBus)},
		PostPingHandler{sp.Get("command.bus").(cqrs.CommandBus)},
	}
}

func (ph *PingHandler) SetEndpoints(r *mux.Router) {
	r.HandleFunc("/ping", ph.getPingHandler.Handle).Methods("GET")
	r.HandleFunc("/ping", ph.postPingHandler.Handle).Methods("POST")
}
