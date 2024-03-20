package main

import (
	"fmt"
	rand2 "math/rand"
	"os"
	"path"
	"strconv"
)

type Entry struct{ Name, Surname, Tel string }

const (
	max = 26
)

var (
	data = []Entry{}
)

func search(key string) []*Entry {
	matches := make([]*Entry, 0)
	for i := range data {
		if data[i].Tel == key {
			matches = append(matches, &data[i])
		}
	}

	return matches
}

func list() {
	for _, entry := range data {
		fmt.Println(entry)
	}
}

func getString(n int) string {
	rs := ""
	for i := 0; i < n; i++ {
		startChar := 'A'
		random := rand2.Intn(max)
		char := string(uint8(startChar) + uint8(random))
		rs += char
	}

	return rs
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		num := strconv.Itoa(rand2.Intn(99) + 100)

		data = append(data, Entry{
			Name:    name,
			Surname: surname,
			Tel:     num,
		})
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	populate(100)
	fmt.Printf("Data has %d entries.\n", len(data))

	// Differentiate between the commands
	switch arguments[1] {
	// The search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		results := search(arguments[2])
		if len(results) == 0 {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		for _, result := range results {
			fmt.Println(*result)
		}
	// The list command
	case "list":
		list()
		// Response to anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
