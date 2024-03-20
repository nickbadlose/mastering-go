package main

import "fmt"

func main() {
	asString := "Hello World! €"
	fmt.Println("First character", string(asString[0]))

	// Runes
	// A Rune
	r := '€'
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as runes
	for _, v := range asString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// Print an existing string as characters
	for _, v := range asString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
