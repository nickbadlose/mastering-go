package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	for _, name := range arguments[1:] {
		match := matchNameSur(name)
		if !match {
			fmt.Printf("%t, %s could not be verified\n", match, name)
			return
		}
	}

	fmt.Println(true)
}
