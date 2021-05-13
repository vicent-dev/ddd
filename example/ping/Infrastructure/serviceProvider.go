package Infrastructure

import (
	"ddd-go/cqrs"
	"ddd-go/example/ping/Application"
	"ddd-go/example/ping/Domain"
	"ddd-go/serviceProvider"
)

func NewServiceProvider() *serviceProvider.ServiceProvider {
	sp := serviceProvider.NewServiceProvider()

	_ = sp.Register("ping.memory.repository", PingInMemoryRepository{})

	//query bus
	pr := sp.Get("ping.memory.repository").(Domain.PingRepository)
	queryBus := cqrs.NewQueryBus()
	queryBus.AddHandler(Application.PingQuery{}, Application.PingQueryHandler{pr})

	//command bus
	commandBus := cqrs.NewCommandBus()
	commandBus.AddHandler(Application.PingCommand{}, Application.PingCommandHandler{pr})

	return sp
}
