package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL := os.Args[1]
	data, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		panic(err)
	}

	data.Body.Close()
}
