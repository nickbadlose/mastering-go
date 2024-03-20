package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["123"] = "456"
	m["key"] = "value"

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	for _, v := range m {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
