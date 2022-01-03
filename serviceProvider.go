package ddd

import (
	"errors"
	"log"
)

type ServiceProvider[T any] struct {
	service map[string]T
}

func NewServiceProvider[T any]() *ServiceProvider[T] {
	return &ServiceProvider[T]{make(map[string]T)}
}

func (sp ServiceProvider[T]) Get(serviceName string) interface{} {

	if _, added := sp.service[serviceName]; !added {
		log.Fatalf("service with name %v not  added", serviceName)
	}

	return sp.service[serviceName]
}

func (sp *ServiceProvider[T]) Register(serviceName string, service T) error {
	if _, added := sp.service[serviceName]; added {
		return errors.New("service with that service name already added")
	}

	sp.service[serviceName] = service
	return nil
}

func (sp *ServiceProvider[T]) OverwriteService(serviceName string, service T) {
	sp.service[serviceName] = service
}
