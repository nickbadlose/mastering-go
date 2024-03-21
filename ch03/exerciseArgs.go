package main

import (
	"fmt"
	"os"
)

type argument struct {
	index int
	value string
}

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("no arguments provided")
	}

	arguments := make([]argument, 0, len(args))
	for i, arg := range args {
		arguments = append(arguments, argument{
			index: i,
			value: arg,
		})
	}

	fmt.Println(arguments)
}
