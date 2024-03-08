package main

import "fmt"

func printAny[T any](s []T) {
	for _, e := range s {
		fmt.Println(e)
	}
}

func print[T string | int](s []T) {
	for _, e := range s {
		fmt.Println(e)
	}
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"One", "Two", "Three"}
	print(Ints)
	print(Strings)
	printAny(Ints)
	printAny(Strings)
}
