package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generatePass(n int64) (string, error) {
	b, err := generateBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	var length int64 = 8
	arguments := os.Args[1:]
	switch len(arguments) {
	case 1:
		t, err := strconv.ParseInt(arguments[0], 10, 64)
		if err == nil {
			length = t
		}
		if length <= 0 {
			length = 8
		}
	default:
		fmt.Println("Using default values!")
	}

	pass, err := generatePass(length)
	if err != nil {
		panic(err)
	}

	fmt.Println(pass[:length])
}
