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
	return MyQueryResult{map[string]string{"": 0}}, nil
}
```
