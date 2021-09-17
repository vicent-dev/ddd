# DDD

Generic tools for making easier to apply DDD with golang

### Run example


```bash
 cd example && go run example
```


### CQRS
#### Command example:

```golang

import (
	"github.com/vicent-dev/ddd"
)

type MyCommand struct{} //command with custom payload

func (c MyCommand) ID() string {
	return "my.command"
}

type MyCommandHandler struct { } //handler with dependencies

func (handler MyCommandHandler) Handle(command ddd.Command) error {
	return nil
}
```
#### Query example:
```golang
import (
	"github.com/vicent-dev/ddd"
)

type MyQuery struct{} //query with custom payload

func (m MyQuery) ID() string {
	return "my.query"
}

//not required sruct but is a good practice, handler function returns interface{}
type MyQueryResult struct { 
	Value map[string]string
}

type MyQueryHandler struct { }//handler with dependencies

func (handler MyQueryHandler) Handle(query ddd.Query) (interface{}, error) {
	return MyQueryResult{map[string]string{"": ""}}, nil
}
```


## Service provider

```golang

import (
	"github.com/vicent-dev/ddd"
)

type MyService struct {}


//init service provider
sp := ddd.NewServiceProvider()

//add services. Maybe a good practice that name of service were a function/property of the same service. We could force implementation of X interface but prefered freedom of services.

sp.Register("function_service", func() string {
    return "my function service"
})

sp.Register("struct_service", MyService{})

//get services where you need it
functionService := sp.Get("function_service")

```

