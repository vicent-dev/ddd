package Infrastructure

import (
	"example/ping/Application"
	"example/ping/Domain"

	"github.com/vicent-dev/ddd"
)

func NewServiceProvider() *ddd.ServiceProvider {
	sp := ddd.NewServiceProvider()

	_ = sp.Register("ping.memory.repository", PingInMemoryRepository{})

	//query bus
	pr := sp.Get("ping.memory.repository").(Domain.PingRepository)
	queryBus := ddd.NewQueryBus()
	queryBus.AddHandler(Application.PingQuery{}, Application.PingQueryHandler{pr})

	sp.Register("query_bus", queryBus)

	//command bus
	commandBus := ddd.NewCommandBus()
	commandBus.AddHandler(Application.PingCommand{}, Application.PingCommandHandler{pr})
	sp.Register("command_bus", commandBus)

	return sp
}
