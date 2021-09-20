package ddd

import (
	"errors"
	"log"
)

type ServiceProvider struct {
	service map[string]interface{}
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{make(map[string]interface{})}
}

func (sp ServiceProvider) Get(serviceName string) interface{} {

	if _, added := sp.service[serviceName]; !added {
		log.Fatalf("service with name %v not  added", serviceName)
	}

	return sp.service[serviceName]
}

func (sp *ServiceProvider) Register(serviceName string, service interface{}) error {
	if _, added := sp.service[serviceName]; added {
		return errors.New("service with that service name already added")
	}

	sp.service[serviceName] = service
	return nil
}

func (sp *ServiceProvider) OverwriteService(serviceName string, service interface{}) {
	sp.service[serviceName] = service
}
