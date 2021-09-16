package ddd

type ChainInterface interface {
	Process(data interface{}) interface{}
}

type Chain struct {
	services []ChainInterface
}

func NewChain(services ...ChainInterface) *Chain {
	return &Chain{services: services}
}

func (c *Chain) Process(data interface{}) interface{} {
	for _, s := range c.services {
		data = s.Process(data)
	}
	return data
}
