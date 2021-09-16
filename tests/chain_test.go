package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type payload struct {
	n1, n2, result float64
}

type sum struct{}

func (s sum) Process(data interface{}) interface{} {
	d := data.(payload)
	d.result = d.n1 + d.n2
	return d
}

type multiplier struct{}

func (m multiplier) Process(data interface{}) interface{} {
	d := data.(payload)
	d.result += d.n1 * d.n2
	return d
}

type subtraction struct{}

func (s subtraction) Process(data interface{}) interface{} {
	d := data.(payload)
	d.result += d.n1 - d.n2
	return d
}

func TestChain_Process(t *testing.T) {
	c := ddd.NewChain(
		sum{},
		multiplier{},
		subtraction{},
	)

	data := c.Process(payload{1, 2, 0}).(payload)
	expected := (1 + 2) + (1 * 2) + (1 - 2)

	if data.result != float64(expected) {
		t.Errorf("Expected %v but get %v", expected, data.result)
	}

}
