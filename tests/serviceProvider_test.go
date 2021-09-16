package tests

import (
	"testing"

	"github.com/vicent-dev/ddd"
)

type formatter struct{}

func (f formatter) format(str string) string {
	return "*** " + str + " ***"
}

func TestServiceProvider_RegisterAndGet(t *testing.T) {

	sp := ddd.NewServiceProvider()
	sp.Register("beautifier", formatter{})
	service := sp.Get("beautifier") //don't return an error, if the service is not loaded, expected log fatal

	if _, ok := service.(formatter); !ok {
		t.Errorf("Expected cast to formatter")
	}
}
