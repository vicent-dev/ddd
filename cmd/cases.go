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
	var queryCase string
	var context string

	if len(params) > 0 {
		context = params[0]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Introduce the name of context: ")
		context, _ = reader.ReadString('\n')
	}

	if len(params) >= 1 {
		queryCase = params[1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Introduce the name of query: ")
		queryCase, _ = reader.ReadString('\n')
	}

	queryCase = strings.TrimSuffix(queryCase, "\n")
	context = strings.TrimSuffix(context, "\n")
	qt := template.NewQueryTemplate(context, queryCase)
	if err := qt.Generate(); err != nil {
		return err
	}

	return nil
}
