package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func main() {
	if len(os.Args) != 2 {
		panic("Usage: file")
	}

	filepath := os.Args[1]

	fileInfo, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}

	if !fileInfo.Mode().IsRegular() {
		panic("not a regular file")
	}

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(f)
	var data []Data
	err = DeSerialize(decoder, &data)
	fmt.Println("After DeSerialize:")
	for index, value := range data {
		fmt.Println(index, value)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err = Serialize(encoder, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("After Serialize:", buf)
}
