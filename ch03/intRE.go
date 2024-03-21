package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchInt(s)
	fmt.Println(ret)

	var t, f int
	for _, pi := range arguments[1:] {
		match := matchInt(pi)
		if !match {
			f++
			fmt.Printf("%t, %s could not be verified\n", match, pi)
			continue
		}
		t++
	}

	fmt.Println("Total true:", t)
	fmt.Println("Total false:", f)
}
