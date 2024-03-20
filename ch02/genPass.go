package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	max           = 94
	defaultLength = 8
)

func randomString(n int) string {
	startChar := '!'
	rs := ""
	for i := 0; i < n; i++ {
		random := rand.Intn(max)
		char := string(uint8(startChar) + byte(random))
		rs = rs + char
	}

	return rs
}

func main() {
	length := defaultLength
	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err == nil {
			length = int(t)
		}
		if length <= 0 {
			length = 8
		}
	default:
		fmt.Println("Using default values...")
	}

	pass := randomString(length)
	fmt.Println(pass)
}
