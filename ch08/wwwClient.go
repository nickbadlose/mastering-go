package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		panic(err)
	}

	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status code:", res.Status)
	header, _ := httputil.DumpResponse(res, false)
	fmt.Print(string(header))

	contentType := res.Header.Get("Content-Type")
	characterSet := strings.Split(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if res.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", res.ContentLength)
	}

	length := 0
	var data []byte
	r := bufio.NewReader(res.Body)
	for {
		n, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + len(n)
		data = append(data, n...)
	}
	fmt.Println("Calculated response data length:", length)

	f, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}
