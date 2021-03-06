package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type menuOption int

const (
	createContextOption menuOption = iota + 1
	createCommandOption
	createQueryOption

	exitString = "q"
)

var (
	menuOptions = map[menuOption]string{
		createContextOption: "Create new context",
		createCommandOption: "Create new use case of type command",
		createQueryOption:   "Create new use case of type query",
	}

	optionsProcessor = map[menuOption]optionCase{
		createContextOption: createContextCase,
		createCommandOption: createCommandCase,
		createQueryOption:   createQueryCase,
	}
)

func menuLoop() {
	reader := bufio.NewReader(os.Stdin)

	var rawOption string

	for rawOption != exitString {
		displayOptions()

		rawOption, _ = reader.ReadString('\n')
		rawOption = strings.TrimSuffix(rawOption, "\n")

		if rawOption == exitString {
			break
		}

		option, err := cleanOption(rawOption)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if err := optionsProcessor[option](os.Args[1:]); err != nil {
			fmt.Println(err.Error())
		}

		break
	}
}

func displayOptions() {
	keys := make([]int, 0, len(menuOptions))

	for k := range menuOptions {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%v) %s \r\n", k, menuOptions[menuOption(k)])
	}

	fmt.Println("Enter q to exit")
}

func cleanOption(rawOption string) (menuOption, error) {
	option, err := strconv.Atoi(rawOption)

	if err != nil {
		return 0, errors.New("Invalid option")
	}

	if _, exist := optionsProcessor[menuOption(option)]; !exist {
		return 0, errors.New("Invalid option")
	}
	return menuOption(option), nil
}
