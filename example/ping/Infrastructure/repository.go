package Infrastructure

import "example/ping/Domain"

type PingInMemoryRepository struct{}

func (r PingInMemoryRepository) Get() Domain.Ping {
	return Domain.CreatePing()
}

func (r PingInMemoryRepository) Create(Domain.Ping) error {
	return nil
}
