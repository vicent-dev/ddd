package main

import "fmt"

type optionCase func(params interface{}) error

var (
	createContextCase = createContext
	createCommandCase = createCommand
	createQueryCase   = createQuery
)

func createContext(params interface{}) error {
	fmt.Println("create context use case")
	return nil
}
func createCommand(params interface{}) error {
	fmt.Println("create command use case")
	return nil
}
func createQuery(params interface{}) error {
	fmt.Println("create query use case")
	return nil
}
