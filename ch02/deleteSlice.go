package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Need one integer value")
		return
	}

	index := args[0]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original slice:", aSlice)

	if len(aSlice) < i {
		fmt.Println("Cannot delete element", i, "out of range")
		return
	}

	aSlice = append(aSlice[:i:len(aSlice)], aSlice[i+1:len(aSlice):len(aSlice)]...)
	fmt.Println("After first deletion:", aSlice)

	if len(aSlice) < i {
		fmt.Println("Cannot delete element", i, "out of range")
		return
	}

	aSlice[i] = aSlice[len(aSlice)-1]
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After second deletion:", aSlice)
}
