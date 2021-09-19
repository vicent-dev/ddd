package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vicent-dev/ddd/cmd/template"
)

type optionCase func(params []string) error

var (
	createContextCase = createContext
	createCommandCase = createCommand
	createQueryCase   = createQuery
)

func createContext(params []string) error {
	var context string
	if len(params) != 0 {
		context = params[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Introduce the name of context: ")
		context, _ = reader.ReadString('\n')
	}

	context = strings.TrimSuffix(context, "\n")
	ct := template.NewContextTemplate(context)
	if err := ct.Generate(); err != nil {
		return err
	}

	return nil
}
func createCommand(params []string) error {
	fmt.Println("create command use case")
	return nil
}
func createQuery(params []string) error {
	fmt.Println("create query use case")
	return nil
}
