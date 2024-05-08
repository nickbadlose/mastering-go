package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type Response struct {
	err         error
	res         *http.Response
	requestTime time.Duration
}

func main() {
	tNow := time.Now()
	if len(os.Args) < 2 {
		panic("Usage: URL")
	}

	if len(os.Args) < 2 {
		panic("Usage: file")
	}

	pflag.Int("n", 10, "Number of requests")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	requests := viper.GetInt("n")

	fmt.Printf("Testing %d requests\n", requests)

	url, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	c := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   15 * time.Second,
	}

	wg := sync.WaitGroup{}
	responseChan := make(chan *Response, requests)
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			t := time.Now()
			defer wg.Done()
			req, rErr := http.NewRequest(http.MethodGet, url.String(), nil)
			if rErr != nil {
				responseChan <- &Response{err: rErr, requestTime: time.Since(t)}
				return
			}
			res, rErr := c.Do(req)
			responseChan <- &Response{res: res, err: rErr, requestTime: time.Since(t)}
		}()
	}

	go func() {
		wg.Wait()
		close(responseChan)
	}()

	var totalBytes []byte
	var failedRequests, successRequests int
	longest := time.Millisecond // Has to be a small value to be able to increase
	shortest := time.Hour       // Has to be a large value to be able to decrease
	for response := range responseChan {
		if response.requestTime > longest {
			longest = response.requestTime
		}
		if response.requestTime < shortest {
			shortest = response.requestTime
		}
		res := response.res
		if response.err != nil {
			failedRequests++
			continue
		}

		if res.StatusCode != http.StatusOK {
			failedRequests++
			continue
		}

		successRequests++
		data, rErr := io.ReadAll(response.res.Body)
		if rErr != nil {
			fmt.Println(rErr)
			continue
		}
		totalBytes = append(totalBytes, data...)
	}

	fmt.Println()
	fmt.Println("Complete requests:", successRequests)
	fmt.Println("Failed requests:", failedRequests)
	fmt.Println("Total transferred:", len(totalBytes))
	fmt.Println("Time taken for tests:", time.Since(tNow))
	fmt.Println("Longest request:", longest)
	fmt.Println("Shortest request:", shortest)
}
