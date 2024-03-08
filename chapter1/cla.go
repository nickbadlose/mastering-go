package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var minArg, maxArg float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		if i == 1 {
			minArg = n
			maxArg = n
			continue
		}

		if n < minArg {
			minArg = n
		}
		if n > maxArg {
			maxArg = n
		}
	}
	fmt.Println("Min:", minArg)
	fmt.Println("Max:", maxArg)
}
